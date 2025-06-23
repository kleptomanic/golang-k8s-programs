// Write a program to delete a pod named `nginx-pod` in the `default` namespace using `client-go`, ensuring graceful termination.

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

	// Build a config file for cluster
	kubeconfigfile, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatalf("Issue on building the config file %v", err)
	}
	// Generating the new config for accessing the cluster
	configset, err := kubernetes.NewForConfig(kubeconfigfile)
	if err != nil {
		log.Fatalf("Issue with building new config %v", err)
	}
	// deleteing the pod having the name nginx
	deletePod := configset.CoreV1().Pods("default").Delete(context.TODO(), "nginx", v1.DeleteOptions{})
	if deletePod != nil {
		log.Fatalf("Issue on deleting the pod %v", deletePod)
	}
	fmt.Println("Successfully deleted the pods.")
}
