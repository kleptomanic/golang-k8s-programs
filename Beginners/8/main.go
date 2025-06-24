// Write a program to delete a deployment named `nginx-deployment` in the `default` namespace using `client-go`.

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
		log.Fatalf("Error building the cofing file : %v", err)

	}
	// Setting the generated config
	clientConfig, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalf("Error setting new config : %v", err)
	}
	deploymentDetails := clientConfig.AppsV1().Deployments("default")
	depStatus := deploymentDetails.Delete(context.TODO(), "nginx-deployment", v1.DeleteOptions{})

	if depStatus != nil {
		log.Fatalf("Failing to delete deployment : %v", depStatus.Error())
	}

	fmt.Println("Deployment deleted successfully")

}
