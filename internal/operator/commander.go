//go:generate mockgen -source=./commander.go -destination=mockCommander.go -package=operator

package operator

import (
	"sync"
)

type (
	Operator interface {
		CreateNewProject(opts *CreateNewProjectOptions) error
		CompressDirectory(opts *CompressDirectoryOptions) error
		ValidateYaml(opts *ValidateYamlOptions) error
	}
	operator struct{}
)

var (
	op   Operator
	once sync.Once
)

func GetInstance() Operator {
	once.Do(func() {
		op = &operator{}
	})
	return op
}
