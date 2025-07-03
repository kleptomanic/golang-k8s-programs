// Write a program to list all ServiceAccounts in the `default` namespace using `client-go`, printing each ServiceAccount’s name.

package main

import (
	"34/connection"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	newConfig, err := connection.CreateNewConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}
	newConnection, err := newConfig.CreateNewConnection()
	if err != nil {
		log.Fatalf("%v", err)
	}
	serviceAccounts, err := newConnection.ClientSet.CoreV1().ServiceAccounts("default").List(newConnection.Ctx, v1.ListOptions{})
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, serviceaccount := range serviceAccounts.Items {
		fmt.Println(serviceaccount.Name)
	}

}
