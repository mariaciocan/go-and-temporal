package main

import (
	"context"
	"fmt"
	"go-and-temporal/app"
	"os"

	"go.temporal.io/sdk/client"
)

func main() {

	c, err := client.NewClient(client.Options{
		HostPort: "127.0.0.1:7233",
	})
	if err != nil {
		fmt.Println("Unable to create client", err)
		os.Exit(1)
	}
	defer c.Close()
	var username string
	var number int
	fmt.Print("Enter your username:")
	fmt.Scanln(&username)
	fmt.Print("Enter a number (0 - 100): ")
	fmt.Scanln(&number)
	workflowOptions := client.StartWorkflowOptions{
		ID:        "user_submission_" + username,
		TaskQueue: "USER_TASK_QUEUE",
	}

	we, err := c.ExecuteWorkflow(
		context.Background(),
		workflowOptions,
		"UserSubmissionWorkFlow",
		app.SubmissionInput{Username: username, Number: number},
	)

	if err != nil {
		fmt.Println("Unable to execute workflow", err)
		os.Exit(1)
	}

	var result app.SubmissionResult
	err = we.Get(context.Background(), &result)
	if err != nil {
		fmt.Println("Workflow failed", err)
		os.Exit(1)
	}

	fmt.Println(result.Message)

}
