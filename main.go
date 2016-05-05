package main

import (
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "gh-token"
	app.Usage = "A simple cli to fetch a github api OAuth token"
	app.Version = version
	app.UsageText = `gh-token [options] <scope> [<scope> ...]

   scope may be one of:
` + commaSeperatedListOfScopes()

	app.Copyright = `gh-token  (c) 2016 Tulip Interfaces
             Provided under terms of Apache License
   go-github (c) 2013 The go-github AUTHORS
                      https://github.com/google/go-github/blob/master/AUTHORS
             Provided under terms of BSD license`

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "note, n",
			Usage: ("create token with `NOTE`.\n" +
				"\tIf not provided will be 'gh-token @ TIMESTAMP'\n"),
		},
		cli.StringFlag{
			Name: "user, u",
			Usage: ("Instead of prompting, acquire token for `USER`\n" +
				"\tAlternatively, use env var:\n\t "),
			EnvVar: "GITHUB_USER",
		},
		cli.StringFlag{
			Name: "password, p",
			Usage: ("Instead of prompting, login with password `PW`\n" +
				"\tAlternatively, use env var:"),
			EnvVar: "GITHUB_PASSWORD",
		},
		cli.StringFlag{
			Name:  "2fa, 2",
			Usage: "2fa `OTP` with which to aquire a token (if needed)",
		},
	}

	app.Action = cmdToken
	app.Run(os.Args)
}

func cmdToken(ctx *cli.Context) (err error) {
	var tokenNote string
	var scopes []github.Scope

	if !ctx.Args().Present() {
		return scopeArgError("You must specify at least one scope")
	}
	for _, scope := range ctx.Args() {
		if v, ok := (*githubScopeMap)[scope]; ok {
			scopes = append(scopes, v)
		} else {
			return scopeArgError("Invalid scope '" + scope + "'")
		}
	}

	if ctx.String("name") == "" {
		tokenNote = "gh-token @ " + time.Now().String()
	}

	transport := github.BasicAuthTransport{
		Username: getUserName(ctx),
		Password: getPassword(ctx),
	}

	if ctx.String("2fa") != "" {
		transport.OTP = ctx.String("2fa")
	}
	var auth *github.Authorization
	for {
		client := github.NewClient(transport.Client())

		auth, _, err = client.Authorizations.Create(&github.AuthorizationRequest{
			Scopes: scopes,
			Note:   &tokenNote,
		})

		if _, ok := err.(*github.TwoFactorAuthError); err != nil && ok {
			transport.OTP = getOTP()
		} else if err != nil {
			return cli.NewExitError(err.Error(), 1)
		} else {
			break
		}
	}

	fmt.Println(*auth.Token)

	return nil
}
