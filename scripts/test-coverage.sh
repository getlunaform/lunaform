#!/usr/bin/env bash

#set -e
#set -x

OUTPUT_PATH="${1}"

echo "mode: atomic" > ${OUTPUT_PATH}/profile.out

non_vendor=$(glide novendor)
all_packages=$(go list ${non_vendor})

for d in ${all_packages}; do

    find ${GOPATH}/src/${d} -maxdepth 0 -type f -name "*.go" -exec false {} \+

    if [ $? -ne 1 ]; then
        go test -v -coverprofile=profile.txt -covermode=atomic $d
        if [ -f profile.txt ]; then
            cat profile.txt | tail -n +2 >> ${OUTPUT_PATH}/profile.out
            rm profile.txt
        fi
    fi

done