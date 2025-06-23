// Write a program to retrieve details of a pod named `nginx-pod` in the `default` namespace using `client-go`, printing its IP address and node name.

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
		log.Fatalf("Error generating files %v", err)
	}
	configSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalf("Error creating new Config %v", err)
	}
	podsName, err := configSet.CoreV1().Pods("default").Get(context.TODO(), "nginx-pod", v1.GetOptions{})

	fmt.Println(podsName.Name, podsName.Spec.NodeName, podsName.Status.PodIP)

}
