// Write a program to list all ConfigMaps in the default namespace using client-go, printing each ConfigMap’s name and data keys.

package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatalf("Error generating config %v", err)
	}
	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalf("Error generating config %v", err)
	}
	coreV1Connection := clientSet.CoreV1()
	configMapsListDetails, err := coreV1Connection.ConfigMaps("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error generating config %v", err)
	}
	for _, j := range configMapsListDetails.Items {
		fmt.Println(j.Name)
		for i, s := range j.Data {
			fmt.Println(i, s)
		}
	}
}
