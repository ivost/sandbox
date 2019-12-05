#!/usr/bin/env bash

APP=hello
#export GO111MODULE=on
#go clean
#go mod tidy
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w"  -o /tmp/${APP} main.go

scp -i ~/.ssh/field_key /tmp/${APP}  brain@192.168.3.1:/tmp

