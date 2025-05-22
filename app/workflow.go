package app

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

func UserSubmissionWorkFlow(ctx workflow.Context, input SubmissionInput) (SubmissionResult, error) {

	var submissions map[string]int
	if workflow.HasLastCompletionResult(ctx) {
		_ = workflow.GetLastCompletionResult(ctx, &submissions)
		// the _ would be of error type, so I could grab it and check if there was an error
	} else {
		submissions = make(map[string]int)
	}
	if _, ok := validUsers[input.Username]; !ok {
		var retryInput SubmissionInput
		workflow.GetLogger(ctx).Info("Invalid username, requesting retry.")
		signalChannel := workflow.GetSignalChannel(ctx, "retry")
		signalChannel.Receive(ctx, &retryInput)
		if _, ok := validUsers[retryInput.Username]; !ok {
			return SubmissionResult{
					false,
					"Invalid username after retry. Failing"},
				nil
		}
		input = retryInput
	}
	if input.Number < 0 || input.Number > 100 {
		return SubmissionResult{
				false,
				"Number must be between 0 and 100. Failing",
			},
			nil
	}

	if _, ok := submissions[input.Username]; ok {
		return SubmissionResult{
				false,
				"You've already added an entry. Failing",
			},
			nil
	}

	submissions[input.Username] = input.Number
	return SubmissionResult{true,
			fmt.Sprintf("Submission accepted for %s:%d", input.Username, input.Number)},
		nil
}
