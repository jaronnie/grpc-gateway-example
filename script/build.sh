#!/bin/bash

set -e

AppName=grpc-gateway-example-app
ROOT_PACKAGE=github.com/jaronnie/grpc-gateway-example
VERSION_PACKAGE="${ROOT_PACKAGE}"/pkg/version

VERSION=$(git describe --dirty --always --tags | sed 's/-/./2' | sed 's/-/./2')
GIT_TAG=$(git describe --tags --abbrev=0)
GIT_COMMIT=$(git log --pretty=format:'%H' -n 1)
GIT_TREE_STATE=$(if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

go mod tidy && \
go build -ldflags "-X ${VERSION_PACKAGE}.gitTag=${GIT_TAG} \
                   -X ${VERSION_PACKAGE}.version=${VERSION} \
                   -X ${VERSION_PACKAGE}.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ') \
                   -X ${VERSION_PACKAGE}.gitCommit=${GIT_COMMIT} \
                   -X ${VERSION_PACKAGE}.gitTreeState=${GIT_TREE_STATE}" \
                   -o cmd/"${AppName}" cmd/main.go