package concourse

import (
	"fmt"
	tasks "github.com/davidwadden/conductor-go-workers/conductor/task"
)

type DeleteConcoursePipelineWorker struct {
}

func NewDeleteConcoursePipelineWorker() *DeleteConcoursePipelineWorker {
	return &DeleteConcoursePipelineWorker{}
}

func (w *DeleteConcoursePipelineWorker) Execute(task *tasks.Task) (*tasks.TaskResult, error) {

	taskResult := tasks.NewTaskResult(task)
	taskResult.Status = tasks.COMPLETED
	taskResult.OutputData["wasDeleted"] = true

	fmt.Println("deleting concourse pipeline")

	return taskResult, nil
}
