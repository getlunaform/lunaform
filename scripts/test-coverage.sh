#!/usr/bin/env bash

OUTPUT_PATH="${1}"
shift;

echo "mode: atomic" > ${OUTPUT_PATH}/profile.out

all_packages=$@

for d in ${all_packages}; do

    if [ $? -ne 1 ]; then
        go test -v -coverprofile=profile.txt -covermode=atomic $d
        if [ -f profile.txt ]; then
            cat profile.txt | tail -n +2 >> ${OUTPUT_PATH}/profile.out
            rm profile.txt
        fi
    fi

done