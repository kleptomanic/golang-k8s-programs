package connection

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Config struct {
	DefaultConfig *rest.Config
}
type Connection struct {
	ClientSet *kubernetes.Clientset
	Ctx       context.Context
}

func CreateNewConfig() (*Config, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &Config{
		DefaultConfig: kubeConfig,
	}, nil
}

func (c Config) CreateNewConnection() (*Connection, error) {
	newConnection, err := kubernetes.NewForConfig(c.DefaultConfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &Connection{
		ClientSet: newConnection,
		Ctx:       context.TODO(),
	}, nil
}
