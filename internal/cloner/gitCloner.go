package cloner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/denizgursoy/gotouch/internal/logger"
	"github.com/denizgursoy/gotouch/internal/store"
)

const (
	GitDirectory = ".git"
)

type (
	gitCloner struct {
		Store  store.Store   `validate:"required"`
		Logger logger.Logger `validate:"required"`
	}
)

func newCloner() Cloner {
	return &gitCloner{
		Store:  store.GetInstance(),
		Logger: logger.NewLogger(),
	}
}

func (g *gitCloner) CloneFromUrl(url, branchName string) error {
	projectFullPath := g.Store.GetValue(store.ProjectFullPath)

	var name plumbing.ReferenceName
	if len(strings.TrimSpace(branchName)) != 0 {
		name = plumbing.NewBranchReferenceName(branchName)
		g.Logger.LogInfo(fmt.Sprintf("Cloning branch %s from   -> %s", branchName, url))
	} else {
		g.Logger.LogInfo("Cloning repository  -> " + url)
	}

	cloneOptions := &git.CloneOptions{
		Depth:         1,
		SingleBranch:  true,
		URL:           url,
		Progress:      os.Stdout,
		ReferenceName: name,
	}

	_, err := git.PlainClone(projectFullPath, false, cloneOptions)
	if err != nil {
		return err
	}

	gitDirectory := filepath.Join(projectFullPath, GitDirectory)
	if err = os.RemoveAll(gitDirectory); err != nil {
		return err
	}
	g.Logger.LogInfo("Cloned successfully")
	return err
}
