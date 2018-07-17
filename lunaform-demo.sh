#!/bin/bash

cd ~/go/src/github.com/drewsonne/lunaform

make

./lunaform-server --port=8080 --scheme=http &

tfs_pid=$!

sleep 4

mkdir -p ~/.config/
cat > ~/.config/lunaform.yaml <<EOF
---
host: localhost
port: 8080
schemes:
  - http
log:
  level: error
apikey: dev-key
EOF

./lunaform terraform workspace create --name live

./lunaform terraform module list

./lunaform terraform module create \
    --name tf-vpc \
    --source github.com/drewsonne/tf-vpc \
    --type git

./lunaform terraform module list
./lunaform terraform stack list
./lunaform terraform stack deploy \
    --name my-vpc \
    --module tf-vpc \
    --workspace live

kill $tfs_pidsla