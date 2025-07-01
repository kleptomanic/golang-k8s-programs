// Write a program to list all namespaces in the cluster using `client-go`, printing each namespace’s name and status.
package main

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type config struct {
	clientconfig *rest.Config
}

type clientConnection struct {
	ctx       context.Context
	clientSet *kubernetes.Clientset
}

func createNewConfig() (*config, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &config{
		clientconfig: kubeConfig,
	}, nil
}

func (c *config) createNewConnection(ctx context.Context) (*clientConnection, error) {
	con, err := kubernetes.NewForConfig(c.clientconfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &clientConnection{
		ctx:       ctx,
		clientSet: con,
	}, nil
}

type namespaceInfo struct {
	Name   string
	Status v1.NamespacePhase
}

type NamespaceInfos []namespaceInfo

func (c clientConnection) getAllNameSpaceList(ctx context.Context) (NamespaceInfos, error) {
	namespaces, err := c.clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	// nses := make([]namespaceInfo,1)
	var nses NamespaceInfos
	// return namespaces.Items, nil
	for _, namespace := range namespaces.Items {
		nses = append(nses, namespaceInfo{
			Name:   namespace.Name,
			Status: namespace.Status.Phase,
		})
	}
	return nses, nil

}

func main() {
	config, err := createNewConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx := context.TODO()
	connection, err := config.createNewConnection(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}
	namespaces, err := connection.getAllNameSpaceList(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, namespace := range namespaces {
		fmt.Printf("Namespace Name : %v (Status: %v)\n", namespace.Name, namespace.Status)
	}

}
