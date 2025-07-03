// Write a program to list all events in the `default` namespace using `client-go`, printing each event’s involved object and message.

package main

import (
	"context"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	e "github.com/kleptomanic/k8s-resouce-generation/events"
	k8scongen "github.com/kleptomanic/k8sConGen"
)

func main() {
	newConfig, err := k8scongen.CreateNewConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx := context.TODO()
	newConnection, err := newConfig.CreateNewConnection(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}
	// events, err := newConnection.ClientSet.CoreV1().Events("default").List(newConnection.Ctx, v1.ListOptions{})
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }
	// for _, event := range events.Items {
	// 	fmt.Println(event.InvolvedObject.Name, event.Message)
	// }
	events, err := e.ListEvent(newConnection.ClientSet, ctx, "default")
	if err != nil {
		log.Fatalf("%v", err)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Object Name", "Message", "Reason"})
	for i, event := range events {
		t.AppendRow([]interface{}{i, event.InvolvedObject.Name, event.Message, event.Reason})
	}
	t.Render()
}
