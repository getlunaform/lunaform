#!/usr/bin/env bash

CWD=$(PWD)


function build_readme() {
    declare -r package=$1
    declare -r file_path=${2}/src/${package}

    count=`(cd ${file_path} && ls -1 *.go 2>/dev/null | wc -l)`
    if [ $count != 0 ]; then
        godoc2md -play -ex ${package} > ${file_path}/README.md
    fi
}

build_readme \
    "github.com/zeebox/terraform-server/server/restapi/controller" \
     ${GOPATH}

build_readme \
    "github.com/zeebox/terraform-server/server/restapi" \
     ${GOPATH}

cd ${CWD}/backend
for package in $(go list ./... | grep -v vendor); do
    build_readme ${package} ${GOPATH}
done

