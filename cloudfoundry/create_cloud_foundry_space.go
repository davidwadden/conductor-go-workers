package cloudfoundry

import (
	"fmt"
	tasks "github.com/davidwadden/conductor-go-workers/conductor/task"
	"code.cloudfoundry.org/cli"
)

type CreateCloudFoundrySpaceWorker struct {

}

func NewCreateCloudFoundrySpaceWorker() *CreateCloudFoundrySpaceWorker {
	return &CreateCloudFoundrySpaceWorker{}
}

func (w *CreateCloudFoundrySpaceWorker) Execute(task *tasks.Task) (*tasks.TaskResult, error) {

	taskResult := tasks.NewTaskResult(task)
	taskResult.Status = tasks.COMPLETED
	taskResult.OutputData["wasDeleted"] = true

	fmt.Println("deleting concourse pipeline")

	return taskResult, nil
}

