package executor

import (
	"errors"
	"log"
	"os"
	"os/exec"

	"github.com/denizgursoy/gotouch/internal/store"
)

type (
	executor struct {
		Store store.Store `validate:"required"`
	}
)

var EmptyRequirementError = errors.New("requirements cannot be empty")

func newExecutor() Executor {
	return &executor{
		Store: store.GetInstance(),
	}
}

func (e executor) Execute(requirements Requirements) error {
	if requirements == nil {
		return EmptyRequirementError
	}

	tasks := make(Tasks, 0)

	for i := 0; i < len(requirements); i++ {
		inputTask, inputRequirements, err := requirements[i].AskForInput()
		if err != nil {
			return err
		}
		tasks = append(tasks, inputTask...)
		requirements = append(requirements, inputRequirements...)
	}

	for _, task := range tasks {
		err := task.Complete()
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *executor) RunCommand(data *CommandData) error {
	if data.WorkingDir == nil {
		projectFullPath := e.Store.GetValue(store.ProjectFullPath)
		err := os.Chdir(projectFullPath)
		if err != nil {
			return err
		}
	}
	cmd := exec.Command(data.Command, data.Args...)

	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return err
	}
	return nil
}
