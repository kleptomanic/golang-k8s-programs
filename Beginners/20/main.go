// Write a program to create a namespace named `test-ns` in the cluster using `client-go`.

package main

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type connection struct {
	clientSet *kubernetes.Clientset
	ctx       context.Context
}

type clientConfig struct {
	config *rest.Config
}

func (c connection) createNameSpace(ns *corev1.Namespace) (*corev1.Namespace, error) {
	namespace, err := c.clientSet.CoreV1().Namespaces().Create(c.ctx, ns, v1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return namespace, nil
}

func createNameSpaceObj() *corev1.Namespace {
	return &corev1.Namespace{
		ObjectMeta: v1.ObjectMeta{
			Name: "test-ns",
		},
	}
}

func (c clientConfig) createNewConnection() (*connection, error) {
	ctx := context.TODO()
	newCon, err := kubernetes.NewForConfig(c.config)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &connection{
		clientSet: newCon,
		ctx:       ctx,
	}, nil

}

func generateConfig() (*clientConfig, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &clientConfig{
		config: kubeConfig,
	}, nil

}

func main() {
	config, err := generateConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}
	newConnection, err := config.createNewConnection()
	if err != nil {
		log.Fatalf("%v", err)
	}
	namespace := createNameSpaceObj()
	createNs, err := newConnection.createNameSpace(namespace)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("Namespace : ", createNs.Name, " is created.")

}
