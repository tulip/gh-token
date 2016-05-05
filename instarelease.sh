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

  /bin/echo -n "Version for this release? " >&2;
  read VERSION
fi

mkdir -p build/Linux build/Darwin
GOOS=linux go build -o build/Linux/$NAME
GOOS=darwin go build -o build/Darwin/$NAME

mkdir -p release
tar -zcf release/${NAME}_${VERSION}_darwin_`uname -m`.tgz -C build/darwin $NAME
tar -zcf release/${NAME}_${VERSION}_linux_`uname -m`.tgz -C build/linux $NAME

git push origin $BRANCH
$GOPATH/bin/gh-release create $REPO $VERSION $BRANCH
