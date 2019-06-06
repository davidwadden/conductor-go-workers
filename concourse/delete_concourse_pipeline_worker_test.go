package concourse_test

import (
	"github.com/davidwadden/conductor-go-workers/concourse"
	tasks "github.com/davidwadden/conductor-go-workers/conductor/task"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeleteConcoursePipelineWorker", func() {

	var worker *concourse.DeleteConcoursePipelineWorker

	BeforeEach(func() {
		worker = concourse.NewDeleteConcoursePipelineWorker()
	})

	It("deletes the concourse pipeline", func() {

		task := tasks.NewTask()
		task.Status = tasks.IN_PROGRESS
		task.InputData = map[string]interface{}{
			"projectName": "some-project-name",
		}

		taskResult, err := worker.Execute(task)

		Expect(err).NotTo(HaveOccurred())
		Expect(taskResult.Status).To(BeEquivalentTo(tasks.COMPLETED))
		Expect(taskResult.OutputData).To(HaveKeyWithValue("wasDeleted", true))
	})

	PIt("does the thing", func() {
		Expect(1).To(Equal(3))
	})
})
