#!/bin/bash
#set -euo pipefail
GO="GO111MODULE=on go"
NAME=mygreet
# dockerhub
DOCKER_REPO=ivostoy
VERSION=0.12.01.0
BRANCH=$(git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown')
BUILD_DATE=$(date +%Y%m%d-%H:%M:%S)
GIT_COMMIT=$(git describe --dirty --always  2> /dev/null  || echo 'unknown')
#BUILDFLAGS="-X main.Version=${VERSION}-X main.Build=${GIT_COMMIT}"
#echo "$BUILDFLAGS"
#go env
set -x
go build -ldflags="-s -w" -o /tmp/server cmd/server/server.go
go build -ldflags="-s -w" -o /tmp/client cmd/client/client.go
set +x
