// Write a program to delete a ConfigMap named `app-config` in the `default` namespace using `client-go`.

package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type clientSet struct {
	clientset *kubernetes.Clientset
}

type configMapDetails struct {
	Name      string
	Namespace string
}

func (c configMapDetails) checkForConfigMapExistance(ctx context.Context, cs *clientSet, cm configMapDetails) error {
	c.Name, c.Namespace = cm.Name, cm.Namespace
	err := cs.clientset.CoreV1().ConfigMaps(c.Namespace).Delete(ctx, c.Name, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func NewClientConnection() (*clientSet, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	clientConn, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &clientSet{clientset: clientConn}, nil
}

func main() {
	newConnection, err := NewClientConnection()
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx := context.TODO()
	configMapData := configMapDetails{
		Name:      "app-config",
		Namespace: "default",
	}
	if err := configMapData.checkForConfigMapExistance(ctx, newConnection, configMapData); err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(configMapData.Name, " is successfully deleted.")

}
