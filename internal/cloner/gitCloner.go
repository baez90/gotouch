package cloner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/denizgursoy/gotouch/internal/auth"
	"github.com/denizgursoy/gotouch/internal/logger"
	"github.com/denizgursoy/gotouch/internal/store"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
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

func (g *gitCloner) CloneFromUrl(rawUrl, branchName string) error {
	projectFullPath := g.Store.GetValue(store.ProjectFullPath)

	var name plumbing.ReferenceName
	if len(strings.TrimSpace(branchName)) != 0 {
		name = plumbing.NewBranchReferenceName(branchName)
		g.Logger.LogInfo(fmt.Sprintf("Cloning branch %s from   -> %s", branchName, rawUrl))
	} else {
		g.Logger.LogInfo("Cloning repository  -> " + rawUrl)
	}

	gitURL, urlParseError := url.Parse(rawUrl)
	if urlParseError != nil {
		return urlParseError
	}

	var gitAuth transport.AuthMethod

	switch gitURL.Scheme {
	case "http", "https":
		gitAuth = auth.NewGitNetrcHTTPAuth()
	}

	cloneOptions := &git.CloneOptions{
		Depth:         1,
		SingleBranch:  true,
		Auth:          gitAuth,
		URL:           rawUrl,
		Progress:      os.Stdout,
		ReferenceName: name,
	}

	_, err := git.PlainClone(projectName, false, cloneOptions)
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
