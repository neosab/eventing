package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/knative/eventing/pkg/apis/eventing/v1alpha1"
	"github.com/knative/eventing/pkg/provisioners"
	provisionerController "github.com/knative/eventing/pkg/provisioners/kafka/controller"
	"github.com/knative/eventing/pkg/provisioners/kafka/controller/channel"
	"github.com/knative/pkg/configmap"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

const (
	BrokerConfigMapKey = "brokers"
)

// SchemeFunc adds types to a Scheme.
type SchemeFunc func(*runtime.Scheme) error

// ProvideFunc adds a controller to a Manager.
type ProvideFunc func(mgr manager.Manager, config *provisionerController.KafkaProvisionerConfig, logger *zap.Logger) (controller.Controller, error)

func main() {
	flag.Parse()
	logf.SetLogger(logf.ZapLogger(false))

	logger := provisioners.NewProvisionerLoggerFromConfig(provisioners.NewLoggingConfig())
	defer logger.Sync()

	// Setup a Manager
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		logger.Error(err, "unable to run controller manager")
		os.Exit(1)
	}

	// Add custom types to this array to get them into the manager's scheme.
	schemeFuncs := []SchemeFunc{
		v1alpha1.AddToScheme,
	}
	for _, schemeFunc := range schemeFuncs {
		schemeFunc(mgr.GetScheme())
	}

	// Add each controller's ProvideController func to this list to have the
	// manager run it.
	providers := []ProvideFunc{
		provisionerController.ProvideController,
		channel.ProvideController,
	}

	provisionerConfig, err := getProvisionerConfig()

	if err != nil {
		logger.Error(err, "unable to run controller manager")
		os.Exit(1)
	}

	for _, provider := range providers {
		if _, err := provider(mgr, provisionerConfig, logger.Desugar()); err != nil {
			logger.Error(err, "unable to run controller manager")
			os.Exit(1)
		}
	}

	mgr.Start(signals.SetupSignalHandler())
}

// getProvisionerConfig returns the details of the associated Provisioner/ClusterChannelProvisioner object
func getProvisionerConfig() (*provisionerController.KafkaProvisionerConfig, error) {
	configMap, err := configmap.Load("/etc/config-provisioner")
	if err != nil {
		return nil, fmt.Errorf("error loading provisioner configuration: %s", err)
	}

	if len(configMap) == 0 {
		return nil, fmt.Errorf("missing provisioner configuration")
	}

	config := &provisionerController.KafkaProvisionerConfig{}

	if value, ok := configMap[BrokerConfigMapKey]; ok {
		brokers := strings.Split(value, ",")
		if len(brokers) == 0 {
			return nil, fmt.Errorf("missing kafka brokers in provisioner configuration")
		}
		config.Brokers = brokers
		return config, nil
	}

	return nil, fmt.Errorf("missing key %s in provisioner configuration", BrokerConfigMapKey)
}
