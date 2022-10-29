//go:generate mockgen -source=./checker.go -destination=mockChecker.go -package=langs

package langs

import (
	"github.com/denizgursoy/gotouch/internal/logger"
	"github.com/denizgursoy/gotouch/internal/store"
	"strings"
)

type Checker interface {
	CheckSetup() error
	CheckDependency(dependency interface{}) error
	GetDependency(dependency interface{}) error
	Setup() error
	CleanUp() error
}

func GetChecker(language string, Logger logger.Logger, str store.Store) Checker {
	if strings.ToLower(language) == "golang" ||
		strings.ToLower(language) == "go" {
		return NewGolangSetupChecker(Logger, str)
	} else {
		return NewEmptySetupChecker()
	}
}
