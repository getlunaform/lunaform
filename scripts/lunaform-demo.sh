#!/bin/bash

cd ~/go/src/github.com/getlunaform/lunaform

make

./lunaform-server --port=8080 --scheme=http --api-key dev-key &

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
    --name terraform-aws-s3 \
    --source github.com/Aplyca/terraform-aws-s3 \
    --type git

./lunaform terraform module list
./lunaform terraform stack list
./lunaform terraform stack deploy \
    --name my-vpc \
    --module terraform-aws-s3 \
    --workspace live \
    --var region=eu-west-1

kill $tfs_pidsla