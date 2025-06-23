// Write a program to list all pods in the default namespace using client-go, printing each pod’s name and status phase to the console.

package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// Load kubeconfig from default location (~/.kube/config)
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	// Build the config from kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating clientset: %v", err)
	}
	// Getting namespace information from the user

	var namespaceName string
	fmt.Print("Enter the Namespace name : ")
	fmt.Scan(&namespaceName)
	// List pods in the "default" namespace
	pods, err := clientset.CoreV1().Pods(namespaceName).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pods: %v", err)
	}

	fmt.Printf("There are %d pods in the %s namespace\n", len(pods.Items), namespaceName)
	for _, pod := range pods.Items {
		fmt.Printf(" - %s: %s\n", pod.Name, pod.Status.Phase)
	}
}
