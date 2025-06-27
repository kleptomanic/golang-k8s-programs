// Write a program to list all Secrets in the `default` namespace using `client-go`, printing each Secret’s name and type.

package main

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type secretDetails struct {
	Name string
	Type string
}

func getSecretDetailsList(k *kubernetes.Clientset, ctx context.Context) []secretDetails {
	secretList, err := k.CoreV1().Secrets("default").List(ctx, v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	var secrets []secretDetails
	for _, secret := range secretList.Items {
		secretType := string(secret.Type)
		secrets = append(secrets, secretDetails{
			Name: secret.Name,
			Type: secretType,
		})
	}
	return secrets
}

func configureConnection() (*kubernetes.Clientset, error) {
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	clientConfig, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return clientConfig, nil
}

func main() {
	ctx := context.TODO()
	newConnection, err := configureConnection()
	if err != nil {
		fmt.Printf("%v", err)
	}
	secretsData := getSecretDetailsList(newConnection, ctx)
	for _, j := range secretsData {
		fmt.Println("Name : ", j.Name, " Type : ", j.Type)
	}

}
