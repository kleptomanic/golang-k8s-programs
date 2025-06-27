// Write a program to create a Secret named `app-secret` in the `default` namespace using `client-go`, with a base64-encoded `password` field.

package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type secretData struct {
	Name       string
	Namespace  string
	SecretData map[string]string
	Type       corev1.SecretType
}

func createNewConnection() (*kubernetes.Clientset, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	clientConnection, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return clientConnection, nil
}

func (s secretData) initilizeSecretData() corev1.Secret {
	return corev1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      s.Name,
			Namespace: s.Namespace,
		},
		StringData: s.SecretData,
		Type:       s.Type,
	}
}
func (s secretData) createNewSecret(k *kubernetes.Clientset, ctx context.Context, c *corev1.Secret) (corev1.Secret, error) {
	secretCreation, err := k.CoreV1().Secrets(s.Namespace).Create(ctx, c, v1.CreateOptions{})
	if err != nil {
		return corev1.Secret{}, fmt.Errorf("%v", err)
	}
	return *secretCreation, nil
}

func main() {
	newConnection, err := createNewConnection()
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.TODO()
	newSecret := secretData{
		Name:      "app-secret",
		Namespace: "default",
		SecretData: map[string]string{
			"password": "test",
		},
		Type: "Opaque",
	}
	secretData := newSecret.initilizeSecretData()
	secretDetails, err := newSecret.createNewSecret(newConnection, ctx, &secretData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(secretDetails.Name, secretData.StringData)

}
