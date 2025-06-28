//  Write a program to list all nodes in the cluster using `client-go`, printing each node’s name and operating system.

package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func createNewConnection() (*kubernetes.Clientset, error) {
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("%v", err)

	}
	clientConnection, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return clientConnection, nil
}

func gatherNodeInformation(k *kubernetes.Clientset, ctx context.Context) (*corev1.NodeList, error) {
	getNodeList, err := k.CoreV1().Nodes().List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return getNodeList, nil
}

func main() {
	newConnection, err := createNewConnection()
	if err != nil {
		fmt.Printf("%v", err)
	}
	ctx := context.TODO()
	getNodeList, err := gatherNodeInformation(newConnection, ctx)
	if err != nil {
		fmt.Printf("%v", err)
	}
	for _, j := range getNodeList.Items {
		fmt.Println(j.Name, j.Status.NodeInfo.OSImage)
	}

}
