# terraform-server
Restful interface to handle terraform deploys. Similiar to the CloudFormation APIs

[![GoDoc](https://godoc.org/github.com/drewsonne/terraform-server/gocd?status.svg)](https://godoc.org/github.com/drewsonne/terraform-server/gocd)
[![Build Status](https://travis-ci.org/drewsonne/terraform-server.svg?branch=master)](https://travis-ci.org/drewsonne/terraform-server)

## Quickstart

```bash
terraform-server --port=8080 --scheme=http

cat > ~/.config/tfs-client.yaml <<EOF
---
host: localhost
port: 8080
schemes:
  - http
log:
  level: error
EOF

tfs-client tf module create \
    --name tf-vpc \
    --source github.com/drewsonne/tf-vpc \
    --type git

tfs-client tf module list
    
tfs-client tf stack deploy \
    --name my-vpc \
    --module tf-vpc

open http://localhost:8080/api/docs


```

## Contributing

This project makes heavy use of the [go-swagger](https://github.com/go-swagger/go-swagger) project. You can install it
on OSX by running:

```bash
brew tap go-swagger/go-swagger
brew install go-swagger
```

For other distributions, see the  [project homepage](https://github.com/go-swagger/go-swagger).
