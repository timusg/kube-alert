package config

import (
	"fmt"

	"github.com/bpineau/kube-alert/pkg/clientset"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// AlertConfig is the main program configuration, passed to controllers Init()
type AlertConfig struct {
	DryRun           bool
	Logger           *logrus.Logger
	ClientSet        *kubernetes.Clientset
	DdAppKey         string
	DdAPIKey         string
	HealthPort       int
	MsgPrefix        string
	SlackToken       string
	SlackChannelName string
}

// Init initialize the configuration (creating the ClientSet for the cluster)
func (c *AlertConfig) Init(apiserver string, kubeconfig string) {
	var err error

	c.ClientSet, err = clientset.NewClientSet(apiserver, kubeconfig)
	if err != nil {
		panic(fmt.Errorf("Failed init Kubernetes clientset: %+v", err))
	}

	_, err = c.ClientSet.Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(fmt.Errorf("Failed to query Kubernetes api-server: %+v", err))
	}

	c.Logger.Info("Kubernetes clientset initialized")
}
