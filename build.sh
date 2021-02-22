#!/usr/bin/env bash
version=$(git describe --tags)
buildTime=$(date '+%Y-%m-%d %H:%M:%S')
echo version:$version
echo buildTime:$buildTime
argsVersion="main.Version=$version"
argsBuildTime="main.BuildTime=$buildTime"
env GOOS=linux GOARCH=arm go build -o myddns -mod=vendor -ldflags="-X '$argsVersion' \
 -X '$argsBuildTime' \"
