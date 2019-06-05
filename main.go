package main

import (
	"fmt"
	"github.com/davidwadden/conductor-go-workers/conductor"
)

func main() {
	conductorClient := conductor.NewConductorHttpClient("http://localhost:8080/api")

	workflowDefs, err := conductorClient.GetAllWorkflowDefs()
	if err != nil {
		panic("amg error")
	}

	fmt.Println("%s", workflowDefs)
}
