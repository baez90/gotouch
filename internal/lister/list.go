package lister

import "sync"

type ProjectStructureData struct {
	Name      string `yaml:"name"`
	Reference string `yaml:"reference"`
	URL       string `yaml:"url"`
}

type Lister interface {
	GetDefaultProjects() ([]*ProjectStructureData, error)
}

var (
	lister Lister
	once   sync.Once
)

func GetInstance() Lister {
	once.Do(func() {
		lister = newGithubLister()
	})
	return lister
}