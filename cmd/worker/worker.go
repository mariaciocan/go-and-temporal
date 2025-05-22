package main

import (
	"go-and-temporal/app"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	c, err := client.NewClient(client.Options{
		HostPort: "127.0.0.1:7233",
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()
	w := worker.New(c, "USER_TASK_QUEUE", worker.Options{})
	w.RegisterWorkflow(app.UserSubmissionWorkFlow)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

}
