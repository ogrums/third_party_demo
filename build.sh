#!/bin/bash
set -ex
export GO111MODULE=on
go env -w GO111MODULE=on

env_version=""
env_version=$1
RUN_DIR=$(cd `dirname $0`; pwd -P)
BUILD_DIR=$RUN_DIR/build
go version
export GOPROXY=https://goproxy.cn
go env -w GOPROXY=https://goproxy.cn

project_name="third_party_server"
if [[ $env_version == "linux" ]];
then

  # for linux
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main_${project_name} main.go
else
  # for mac
  go build -o bin/main_${project_name} main.go
fi

./bin/main_third_party_server
