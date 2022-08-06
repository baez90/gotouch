package extractor

import (
	"fmt"
	"github.com/artdarek/go-unzip"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type gzipExtractor struct {
}

func newGzipExtractor() Extractor {
	return gzipExtractor{}
}

func (g gzipExtractor) Extract(url, projectName string) {
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		println(err)
		return
	}
	filePath2 := filepath.Join(os.TempDir(), filepath.Base(url))

	create, err := os.Create(filePath2)
	_, err = io.Copy(create, response.Body)

	if err != nil {
		fmt.Println("error")
	}
	target := fmt.Sprintf("./%s", projectName)
	uz := unzip.New(filePath2, target)

	err = uz.Extract()
	if err != nil {
		fmt.Println(err)
	}
}
