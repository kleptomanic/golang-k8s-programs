// Write a program to list all deployments in the `default` namespace using `client-go`, printing each deployment’s name and replica count.

package main

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatalf("Error Building Config File : %v", err)
	}
	configSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalf("Error generating new files : %v", err)
	}
	deployment, err := configSet.AppsV1().Deployments("default").List(context.TODO(), v1.ListOptions{})

	if err != nil {
		log.Fatalf("Issue reading the deployments: %v", err)
	}

	fmt.Printf("%d deployments found in default Namespace.\n", len(deployment.Items))
	for _, dep := range deployment.Items {
		fmt.Println(dep.Name, int(*dep.Spec.Replicas))
	}
}
