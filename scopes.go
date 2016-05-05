package main

import (
	"bytes"
	"fmt"
	"syscall"
	"text/tabwriter"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
)

// there doesn't appear to me to be a better way to fetch this from the
// underlying library than just to copy it from the source. Alas.
// https://github.com/google/go-github/blob/master/github/authorizations.go
var allGithubScopes = []github.Scope{
	github.ScopeUser,
	github.ScopeUserEmail,
	github.ScopeUserFollow,
	github.ScopePublicRepo,
	github.ScopeRepo,
	github.ScopeRepoDeployment,
	github.ScopeRepoStatus,
	github.ScopeDeleteRepo,
	github.ScopeNotifications,
	github.ScopeGist,
	github.ScopeReadRepoHook,
	github.ScopeWriteRepoHook,
	github.ScopeAdminRepoHook,
	github.ScopeAdminOrgHook,
	github.ScopeReadOrg,
	github.ScopeWriteOrg,
	github.ScopeAdminOrg,
	github.ScopeReadPublicKey,
	github.ScopeWritePublicKey,
	github.ScopeAdminPublicKey,
}

func scopeArgError(message string) cli.ExitCoder {
	buf := new(bytes.Buffer)
	tabw := new(tabwriter.Writer)
	tabw.Init(buf, 0, 8, 1, '\t', 0)
	fmt.Fprint(tabw, commaSeperatedListOfScopes())
	tabw.Flush()
	return cli.NewExitError(
		message+"\nValid scopes are"+buf.String(),
		int(syscall.EINVAL),
	)
}

var githubScopeMap = (func() *map[string]github.Scope {
	s := make(map[string]github.Scope, len(allGithubScopes))
	for _, v := range allGithubScopes {
		s[string(v)] = v
	}
	return &s
})()

func commaSeperatedListOfScopes() string {
	var buff bytes.Buffer
	i := 0
	buff.WriteString("     ")
	for _, scope := range allGithubScopes {
		next := fmt.Sprintf("%s\t", scope)
		if i%4 == 0 {
			buff.WriteString("\n     ")
		}
		buff.WriteString(next)
		i++
	}
	return buff.String()
}
