NAME:=$(shell basename $$PWD)
ARCH:=$(shell uname -m)
REPO:=$(shell git config --get remote.origin.url | perl -ne 'm{github.com[:/](.+/[^.]+)}; print  $$1')
VERSION=0.0.1

build:
	go build -ldflags '-X main.version=$(VERSION)-dev'

build_all:
	mkdir -p build/linux && GOOS=linux go build -ldflags '-w -s -X main.version=$(VERSION)' -o build/linux/$(NAME)
	mkdir -p build/darwin && GOOS=darwin go build -ldflags '-w -s -X main.version=$(VERSION)' -o build/darwin/$(NAME)

release: build_all
	upx build/linux/$(NAME)
	upx build/darwin/$(NAME)
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_darwin_$(ARCH).tgz -C build/darwin $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_linux_$(ARCH).tgz -C build/linux $(NAME)
	gh-release create $(REPO) $(VERSION) master

.PHONY: release build build_all
