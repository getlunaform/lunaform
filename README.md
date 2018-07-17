![lunaform](./docs/lunaform.jpeg)

Restful interface to handle terraform deploys. Similiar to the CloudFormation APIs

[![GoDoc](https://godoc.org/github.com/drewsonne/lunaform/gocd?status.svg)](https://godoc.org/github.com/drewsonne/lunaform/gocd)
[![Build Status](https://travis-ci.org/drewsonne/lunaform.svg?branch=master)](https://travis-ci.org/drewsonne/lunaform)

## Quickstart

```bash
brew tap drewsonne/tap
brew install lunaform

lunaform-server \
    --port=8080 \
    --scheme=http \
    --api-key=dev-key

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

lunaform terraform module create \
    --name tf-vpc \
    --source github.com/drewsonne/tf-vpc \
    --type git
lunaform terraform module list

lunaform terraform workspace create \
    --name live
lunaform terraform workspace list

    
lunaform terraform stack deploy \
    --name my-vpc \
    --module tf-vpc \
    --workspace live

open http://localhost:8080/api/docs


```

## Contributing

This project makes heavy use of the [go-swagger](https://github.com/go-swagger/go-swagger) project. You can install it
on OSX by running:

```bash
brew tap go-swagger/go-swagger
brew install go-swagger
```

For dependency management we use glide:

```bash
brew install glide
glide install
```

For other distributions, see the  [project homepage](https://github.com/go-swagger/go-swagger).
