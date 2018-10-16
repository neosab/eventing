/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package channel

import (
	"fmt"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/knative/eventing/pkg/apis/eventing/v1alpha1"
	common "github.com/knative/eventing/pkg/provisioners/kafka/controller"
)

const (
	// controllerAgentName is the string used by this controller to identify
	// itself when creating events.
	controllerAgentName = "kafka-provisioner-channel-controller"
)

type reconciler struct {
	client            client.Client
	recorder          record.EventRecorder
	logger            *zap.Logger
	config            *common.KafkaProvisionerConfig
	kafkaClusterAdmin sarama.ClusterAdmin
}

// Verify the struct implements reconcile.Reconciler
var _ reconcile.Reconciler = &reconciler{}

// ProvideController returns a Channel controller.
func ProvideController(mgr manager.Manager, config *common.KafkaProvisionerConfig, logger *zap.Logger) (controller.Controller, error) {
	// Setup a new controller to Reconcile Channel.
	clusterAdmin, err := getKafkaAdminClient(config)
	if err != nil {
		return nil, fmt.Errorf("unable to build kafka admin client: %s", err)
	}
	c, err := controller.New(controllerAgentName, mgr, controller.Options{
		Reconciler: &reconciler{
			recorder:          mgr.GetRecorder(controllerAgentName),
			logger:            logger,
			config:            config,
			kafkaClusterAdmin: clusterAdmin,
		},
	})
	if err != nil {
		return nil, err
	}

	// Watch Channel events and enqueue Channel object key.
	if err := c.Watch(&source.Kind{Type: &v1alpha1.Channel{}}, &handler.EnqueueRequestForObject{}); err != nil {
		return nil, err
	}

	return c, nil
}

func getKafkaAdminClient(config *common.KafkaProvisionerConfig) (sarama.ClusterAdmin, error) {
	saramaConf := sarama.NewConfig()
	saramaConf.Version = sarama.V1_1_0_0
	saramaConf.ClientID = controllerAgentName
	return sarama.NewClusterAdmin(config.Brokers, saramaConf)
}

func (r *reconciler) InjectClient(c client.Client) error {
	r.client = c
	return nil
}
