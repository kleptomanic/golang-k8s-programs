// Write a program to list all services in the `default` namespace using `client-go`, printing each service’s name and cluster IP.

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
		log.Fatalf("Error building the configuration : %v", err)
	}
	clientConfig, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalf("Error building new configuration : %v", err)
	}
	serviceInfoHolder := clientConfig.CoreV1().Services("default")
	serviceLists, err := serviceInfoHolder.List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Fatalf("Error Listing services : %v", err)
	}
	for _, j := range serviceLists.Items {
		fmt.Println(j.Name, j.Spec.ClusterIPs)

		for _, i := range j.Spec.Ports {
			fmt.Println(i.Name, i.TargetPort, i.Port, i.NodePort)
		}
	}
}
