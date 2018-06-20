#!/usr/bin/env bash

declare -ar folders=( "backend" "server" )

ROOT_DIR=$(PWD)


function build_readme() {
    declare -r package=$1
    declare -r file_path=${2}/src/${package}

    count=`(cd ${file_path} && ls -1 *.go 2>/dev/null | wc -l)`
    if [ $count != 0 ]; then
        godoc2md -play -ex ${package} > ${file_path}/README.md
    fi
}

build_readme \
    "github.com/drewsonne/terraform-server/server/restapi" \
     ${GOPATH}

for folder in "${folders[@]}"; do
    echo "${folder}"
    cd ${ROOT_DIR}/${folder}
    for package in $(go list ./... | grep -v vendor); do
        build_readme ${package} ${GOPATH}
    done
done

