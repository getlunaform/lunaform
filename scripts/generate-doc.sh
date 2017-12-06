#!/usr/bin/env bash

CWD=$(PWD)


function build_readme() {
    declare -r package=$1
    declare -r gopath=$2
    godoc2md ${package} > ${gopath}/src/${package}/README.md
}

cd ${CWD}/backend
for package in $(go list ./... | grep -v vendor); do
    build_readme ${package} ${GOPATH}
done

build_readme \
    "github.com/zeebox/terraform-server/server/restapi/controller" \
     ${GOPATH}

build_readme \
    "github.com/zeebox/terraform-server/server/restapi" \
     ${GOPATH}