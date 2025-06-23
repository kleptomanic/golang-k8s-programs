// Write a program to create a pod named `nginx-pod` in the `default` namespace using `client-go`, running an `nginx` container with the latest image.

package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	config := filepath.Join(homedir.HomeDir(), ".kube", "config")

	// Generating config from the default config path

	kubeconfig, err := clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		log.Fatalf("Error generating the config %v", err)
	}
	configset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		log.Fatalf("Issue Generating the config %v", err)
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx",
					Image: "nginx&latest",
				},
			},
		},
	}
	createPod, err := configset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error creating the pod %v", err)
	}
	fmt.Printf("Pod %s created in namespace %s\n", createPod.Name, createPod.Namespace)
}
