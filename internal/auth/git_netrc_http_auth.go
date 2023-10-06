package auth

import (
	"fmt"
	"net/http"

	httptransport "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/jdx/go-netrc"
)

var _ httptransport.AuthMethod = (*GitNetrcHTTPAuth)(nil)

func NewGitNetrcHTTPAuth() *GitNetrcHTTPAuth {
	creds, err := ParseNetrc()
	if err != nil {
		return new(GitNetrcHTTPAuth)
	}

	return &GitNetrcHTTPAuth{
		credentials: creds,
	}
}

type GitNetrcHTTPAuth struct {
	credentials *netrc.Netrc
}

func (g GitNetrcHTTPAuth) String() string {
	if g.credentials == nil {
		return "netrc://<empty>"
	}

	return fmt.Sprintf("netrc://%s", g.credentials.Path)
}

func (GitNetrcHTTPAuth) Name() string {
	return "http-netrc-auth"
}

func (g GitNetrcHTTPAuth) SetAuth(r *http.Request) {
	if g.credentials == nil {
		return
	}

	machine := g.credentials.Machine(r.Host)
	if machine == nil {
		return
	}

	r.SetBasicAuth(machine.Get("login"), machine.Get("password"))
}
