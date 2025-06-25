// Write a program to update a ConfigMap named `app-config` in the `default` namespace using `client-go`, changing the value of `key` to `new-value`.

package main

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type cmAppConfig struct {
	Name      string
	Namespace string
	DataKeys  map[string]string
}

func (c *cmAppConfig) configMapAppData() *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.Name,
			Namespace: c.Namespace,
		},
		Data: c.DataKeys,
	}
}

func (c cmAppConfig) checkingConfigMapExistance(ctx context.Context, nc *newConnection) (*corev1.ConfigMap, error) {
	check, err := nc.clientSet.CoreV1().ConfigMaps(c.Namespace).Get(ctx, c.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return check, nil
}

func (c cmAppConfig) updateConfigMap(ctx context.Context, nc *newConnection, cm *corev1.ConfigMap) (*corev1.ConfigMap, error) {
	update, err := nc.clientSet.CoreV1().ConfigMaps(c.Namespace).Update(ctx, cm, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return update, nil
}

type newConnection struct {
	clientSet *kubernetes.Clientset
}

func createNewConnection() (*newConnection, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	clientConnection, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &newConnection{clientSet: clientConnection}, nil

}

func main() {
	ctx := context.TODO()
	connectionConfig, err := createNewConnection()
	if err != nil {
		log.Fatalf("%v", err)
	}
	newCmAppConfig := cmAppConfig{
		Name:      "app-config",
		Namespace: "default",
		DataKeys: map[string]string{
			"keys": "new-value",
		},
	}
	updatedCmConfig := newCmAppConfig.configMapAppData()
	cmAppData, err := newCmAppConfig.checkingConfigMapExistance(ctx, connectionConfig)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("Config Map found.", cmAppData.Name)
	cmAppDataUpdate, err := newCmAppConfig.updateConfigMap(ctx, connectionConfig, updatedCmConfig)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(cmAppDataUpdate.Name, cmAppDataUpdate.Data)

}
