// Write a program to create a ConfigMap named `app-config` in the `default` namespace using `client-go`, with a key-value pair `key=value`.

package main

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type configMapStruct struct {
	Name      string
	Namespace string
	DataKeys  map[string]string
}

func (c *configMapStruct) generateCmObject() *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: v1.ObjectMeta{
			Name:      c.Name,
			Namespace: c.Namespace,
		},
		Data: c.DataKeys,
	}
}

func (cc clientConnection) createNewConfigMap(ctx context.Context, acm *configMapStruct) error {
	gcm := acm.generateCmObject()

	_, err := cc.ClientSet.CoreV1().ConfigMaps(acm.Namespace).Create(ctx, gcm, v1.CreateOptions{})
	return err
}

type clientConnection struct {
	ClientSet *kubernetes.Clientset
}

func createConnection() (*clientConnection, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &clientConnection{ClientSet: clientSet}, nil
}

func main() {
	newConnection, err := createConnection()
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx := context.TODO()
	configMapSchema := &configMapStruct{
		Name:      "app-config",
		Namespace: "default",
		DataKeys: map[string]string{
			"keys": "data",
		},
	}
	if err := newConnection.createNewConfigMap(ctx, configMapSchema); err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("Config Map is created : ", configMapSchema.Name)
}
