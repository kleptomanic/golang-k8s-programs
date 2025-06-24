// Write a program to create a ClusterIP service named `nginx-service` in the `default` namespace using `client-go`, exposing port 80 for pods labeled `app=nginx`.

package main

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatalf("Error building configuration %v", err)
	}
	clientConfig, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalf("Error generating new config %v", err)
	}
	serviceSpec := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx-service",
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeClusterIP,
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Protocol:   corev1.ProtocolTCP,
					Port:       80,
					TargetPort: intstr.FromInt(80),
				},
			},
			Selector: map[string]string{"app": "nginx"},
		},
	}
	podSpec := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "nginx-pod",
			Labels: map[string]string{"app": "nginx"},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx-container",
					Image: "nginx",
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 80,
						},
					},
				},
			},
		},
	}
	linkStatus := clientConfig.CoreV1()
	podGenerate, err := linkStatus.Pods("default").Create(context.TODO(), &podSpec, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Issue Creating the Pod : %v", err)
	}
	fmt.Println("Pod ", podGenerate.Name, " is created.")
	serviceCreate, err := linkStatus.Services("default").Create(context.TODO(), &serviceSpec, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Issue Creating the Service : %v", err)
	}
	fmt.Println("Pod ", serviceCreate.Name, " is created.")

}
