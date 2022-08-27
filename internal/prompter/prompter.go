//go:generate mockgen -source=./prompter.go -destination=mockPrompter.go -package=prompter

package prompter

import (
	"errors"
	"github.com/denizgursoy/gotouch/internal/manager"
	"sync"
)

var (
	once                           = sync.Once{}
	prompter                       Prompter
	ErrProductStructureListIsEmpty = errors.New("options can not be empty")
)

type (
	Prompter interface {
		AskForString(direction string, validator StringValidator) (string, error)
		AskForSelectionFromList(direction string, list []Option) (interface{}, error)
	}

	ListOption struct {
		DisplayText string
		ReturnVal   interface{}
	}

	StringValidator func(string) error

	Option interface {
		String() string
	}
)

func GetInstance() Prompter {
	once.Do(func() {
		prompter = &promptUi{
			Manager: manager.GetInstance(),
		}
	})
	return prompter
}