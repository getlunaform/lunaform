#!/bin/bash

cd ~/go/src/github.com/drewsonne/lunarform

./lunarform --port=8080 --scheme=http &

tfs_pid=$!

sleep 4

cat > ~/.config/tfs-client.yaml <<EOF
---
host: localhost
port: 8080
schemes:
  - http
log:
  level: error
EOF

./tfs-client tf module list

./tfs-client tf module create \
    --name tf-vpc \
    --source github.com/drewsonne/tf-vpc \
    --type git

./tfs-client tf module list
./tfs-client tf stack list
./tfs-client tf stack deploy \
    --name my-vpc \
    --module tf-vpc

kill $tfs_pid