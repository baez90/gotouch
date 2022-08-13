package lister

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

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
		lister = newHttpLister(&http.Client{}, &PropertiesUrl)
	})
	return lister
}

func ParseToProjectStructureData(reader io.ReadCloser) ([]*ProjectStructureData, error) {
	data := make([]*ProjectStructureData, 0)

	allBytes, err := ioutil.ReadAll(reader)
	err = yaml.Unmarshal(allBytes, &data)

	if err != nil {
		log.Println(err.Error())
		return nil, ProjectDataParseError
	}

	if data == nil {
		return nil, NoProjectError
	}

	return data, nil
}
