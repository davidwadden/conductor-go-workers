package main

import (
	"fmt"
	"github.com/davidwadden/conductor-go-workers/concourse"
	"github.com/davidwadden/conductor-go-workers/conductor"
	"github.com/davidwadden/conductor-go-workers/conductor/task/sample"
)

const baseUrl = "http://localhost:8080/api"



func main() {
	conductorClient := conductor.NewConductorHttpClient(baseUrl)

	workflowDefs, err := conductorClient.GetAllWorkflowDefs()
	if err != nil {
		panic("amg error")
	}

	fmt.Println("%s", workflowDefs)

	conductorWorker := conductor.NewConductorWorker(baseUrl, 1, 1000)
	conductorWorker.Start("create_jira_project", sample.Task_1_Execution_Function, false)
	conductorWorker.Start("task_2", sample.Task_2_Execution_Function, false)

	concourseWorker := concourse.NewDeleteConcoursePipelineWorker()
	conductorWorker.Start("delete_concourse_pipeline", concourseWorker.Execute, false)

	select {}
}
