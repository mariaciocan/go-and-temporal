package app

type SubmissionInput struct {
	Username string
	Number   int
}

type SubmissionResult struct {
	Success bool
	Message string
}
