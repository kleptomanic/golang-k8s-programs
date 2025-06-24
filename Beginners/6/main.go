// Write a program to create a deployment named `nginx-deployment` in the `default` namespace using `client-go`, with 3 replicas of an `nginx` container.

package main

import (
	"context"
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)

	if err != nil {
		log.Fatalf("Issue building the configuration : %v", err)
	}

	configSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalf("Issue building the configuration : %v", err)
	}
	replicas := int32(3)
	deployment := &appsv1.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name:      "nginx-deployment",
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &v1.LabelSelector{
				MatchLabels: map[string]string{"host": "dev"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: map[string]string{"host": "dev"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx",
						},
					},
				},
			},
		},
	}

	createdDeployment, err := configSet.AppsV1().Deployments("default").Create(context.TODO(), deployment, v1.CreateOptions{})

	if err != nil {
		log.Fatalf("Issue creating the deployment %v", err)
	}
	fmt.Println("Deployment is created.", createdDeployment.Status)
}
