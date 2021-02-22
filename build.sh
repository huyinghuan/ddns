#!/usr/bin/env bash
version=$(git describe --tags)
buildTime=$(date '+%Y-%m-%d %H:%M:%S')
echo version:$version
echo buildTime:$buildTime
argsVersion="main.Version=$version"
argsBuildTime="main.BuildTime=$buildTime"
env GOOS=linux GOARCH=arm go build -o out/myddns_arm -mod=vendor -ldflags="-X '$argsVersion' \
 -X '$argsBuildTime'"

env GOOS=linux GOARCH=arm64 go build -o out/myddns_arm64 -mod=vendor -ldflags="-X '$argsVersion' \
 -X '$argsBuildTime' "

 env GOOS=linux GOARCH=amd64 go build -o out/myddns_amd64 -mod=vendor -ldflags="-X '$argsVersion' \
 -X '$argsBuildTime' "