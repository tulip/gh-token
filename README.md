# gh-token
A simple cli to fetch a github api OAuth token

Inspired by the likes of  [progrium/gh-release](https://github.com/progrium/gh-release) and [icholy/gist](https://github.com/icholy/gist),

 And noticing that, while these sorts of utilities required an oauth token from the environment, they wanted you to \**gasp*\* use a browser to make it,
<sub>(And wanting to practice writing more Go...)</sub>

I wrote this simple utility to make Github personal access tokens

## Installation

```
go get github.com/tulip/gh-token
```

or Download it from the [Releases](https://github.com/tulip/gh-token/releases)

## Example uses

`auth_and_gist.sh`
```bash
#!/bin/sh
go get github.com/icholy/gist
GITHUB_TOKEN=${GITHUB_TOKEN-`gh-token gist`} $GOPATH/bin/gist
```

`instarelease.sh` - a zero-to-released interactive two platform go binary github release script
```sh
#!/bin/sh
go get github.com/progrium/gh-release
go get github.com/tulip/gh-token

# get a new token if none set in environment
GITHUB_ACCESS_TOKEN=${GITHUB_TOKEN-`$GOPATH/bin/gh-token repo`}

NAME=${NAME-`basename $PWD`}
REPO=${REPO-`git config --get remote.origin.url | perl -ne 'm{github.com[:/](.+/[^.]+)}; print  $1'`}
BRANCH=${BRANCH-`git rev-parse --abbrev-ref`}
if [ -z $VERSION ]; then
  git fetch --tags
  if git describe --tags --match='v*' --abbrev=0 &>/dev/null; then
    echo "Last version was `git describe --tags --match='v*' --abbrev=0`"
  fi

  /bin/echo -n "Version for this release? ";
  read VERSION
fi

echo ""
mkdir -p build/Linux build/Darwin
echo "Building for Linux... "
GOOS=linux go build -o build/Linux/$NAME
echo "Building for Darwin... "
GOOS=darwin go build -o build/Darwin/$NAME

echo "Compressing..."
mkdir -p release
tar -zcf release/${NAME}_${VERSION}_darwin_`uname -m`.tgz -C build/darwin $NAME
tar -zcf release/${NAME}_${VERSION}_linux_`uname -m`.tgz -C build/linux $NAME

git push origin $BRANCH
$GOPATH/bin/gh-release create $REPO $VERSION $BRANCH
```

I could also imagine it being useful for the entrypoint script of certain kinds of docker containers?

## Usage

available as a `--help` command, reproduced here for your convenience

```
NAME:
   gh-token - A simple cli to fetch a github api OAuth token

USAGE:
   gh-token [options] <scope> [<scope> ...]

   scope may be one of:

     user             user:email      user:follow       public_repo   
     repo             repo_deployment repo:status       delete_repo   
     notifications    gist            read:repo_hook    write:repo_hook   
     admin:repo_hook  admin:org_hook  read:org          write:org   
     admin:org        read:public_key write:public_key  admin:public_key  

VERSION:
   0.0.1

GLOBAL OPTIONS:
   --note NOTE, -n NOTE   create token with NOTE.
                          If not provided will be 'gh-token @ TIMESTAMP'
   --user USER, -u USER   Instead of prompting, acquire token for USER
                          Alternatively, use env var: [$GITHUB_USER]
   --password PW, -p PW   Instead of prompting, login with password PW
                          Alternatively, use env var: [$GITHUB_PASSWORD]
   --2fa OTP, -2 OTP      2fa OTP with which to aquire a token (if needed)
   --help, -h             show help
   --version, -v          print the version

COPYRIGHT:
   gh-token  (c) 2016 Tulip Interfaces
             Provided under terms of Apache License
   go-github (c) 2013 The go-github AUTHORS
                      https://github.com/google/go-github/blob/master/AUTHORS
             Provided under terms of BSD license


```

## Contributing

How to submit changes:

1. Fork this repository.
2. Make your changes.
3. Email us at opensource@tulip.co to sign a CLA.
4. Submit a pull request.


## Who's Behind It

gh-token is maintained by Tulip. We're an MIT startup located in Boston, helping enterprises manage, understand, and improve their manufacturing operations. We bring our customers modern web-native user experiences to the challenging world of manufacturing, currently dominated by ancient enterprise IT technology. We work on Meteor web apps, embedded software, computer vision, and anything else we can use to introduce digital transformation to the world of manufacturing. If these sound like interesting problems to you, [we should talk](mailto:jobs@tulip.co).


## License

gh-token is licensed under the [Apache Public License](LICENSE).
