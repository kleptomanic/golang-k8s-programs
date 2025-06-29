// Write a program to retrieve details of a node named `node-1` using `client-go`, printing its allocatable CPU and storage capacity.

package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type clientConfig struct {
	newConfig *kubernetes.Clientset
}

func (c *clientConfig) getNodeInformation(ctx context.Context, node string) (*corev1.Node, error) {
	return c.newConfig.CoreV1().Nodes().Get(ctx, node, v1.GetOptions{})
}
func (c *clientConfig) createNewConfig() (*kubernetes.Clientset, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	newClientConfig, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	c.newConfig = newClientConfig
	return newClientConfig, nil
}

func main() {
	clientSet := &clientConfig{}
	_, err := clientSet.createNewConfig()
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.TODO()
	nodeDetails, err := clientSet.getNodeInformation(ctx, "minikube")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nodeDetails.Status.Capacity.Cpu(), nodeDetails.Status.Capacity.Storage(), nodeDetails.Status.Capacity.Memory())

}
