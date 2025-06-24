// Write a program to update the replica count of a deployment named `nginx-deployment` to 5 in the `default` namespace using `client-go`.

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
	// Creating a config file and setting the config to perform requests

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)

	if err != nil {
		log.Fatalf("Issue building the config : %v", err)
	}
	configSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalf("Error generating new config : %v", err)
	}

	getSpecificDeployment, err := configSet.AppsV1().Deployments("default").Get(context.TODO(), "nginx-deployment", metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error getting the deployments status %v", err)
	}

	newReplicas := int32(5)

	if int(*getSpecificDeployment.Spec.Replicas) <= 3 {
		getSpecificDeployment.Spec.Replicas = &newReplicas

		updatedDeployment, err := configSet.AppsV1().Deployments("default").Update(context.TODO(), getSpecificDeployment, metav1.UpdateOptions{})

		if err != nil {
			log.Fatalf("Issue updating the deployments : %v", err)
		}
		fmt.Printf("Deployment %v is updated to %v replicas", updatedDeployment.Name, int(*updatedDeployment.Spec.Replicas))
	}

}
