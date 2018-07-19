.DEFAULT_GOAL := default

SRC_YAML?=${CURDIR}"/swagger.yml"
CGO?=cgo

CWD?=$(shell pwd)

SHELL:=/bin/bash
ENVIRONMENT?=DEVELOPMENT

BUILD_ID?=$(ENVIRONMENT)

GOROOT?=${HOME}/go

GO_TARGETS= ./server ./backend
GOR_TARGETS= ./server/... ./backend/...

VERSION?=$(shell git rev-parse --short HEAD)

PARENT_DIR=$(dir ${CURDIR})
CUR_DIR_NAME=$(notdir ${CURDIR})

MODEL_PACKAGE?=${GOROOT}/src/github.com/getlunaform/lunaform-models-go
CLIENT_PACKAGE?=${GOROOT}/src/github.com/getlunaform/lunaform-client-go
JS_CLIENT_PACKAGE?=${GOROOT}/src/github.com/getlunaform/lunaform-client-js

##################
# Global Targets #
##################
build: build-server build-client-go
clean: clean-server clean-client-go
generate: generate-server generate-client-go

default: clean generate build

update-vendor:
	glide update

##################
# Server targets #
##################

build-server:
	go build \
		-a -installsuffix $(CGO) \
		-o $(CWD)/lunaform-server \
		github.com/getlunaform/lunaform/server/cmd/lunaform-server

clean-server:
	rm -rf $(CWD)/server/cmd/ \
		$(CWD)/server/restapi/operations \
		$(CWD)/server/restapi/doc.go \
		$(CWD)/server/restapi/embedded_spec.go \
		$(CWD)/server/restapi/server.go \
		$(CWD)/lunaform \
		$(CWD)/profile.txt

generate-server:
	swagger generate server \
		--target=server \
		--principal=models.ResourceAuthUser \
		--name=lunaform \
		--existing-models=github.com/getlunaform/lunaform-models-go \
		--skip-models \
		--spec=$(SRC_YAML)

run-server:
	$(CWD)/lunaform --port=8080 --scheme=http

##################
# Client targets #
##################
clean-client-go:
	$(MAKE) -C $(CLIENT_PACKAGE) clean

generate-client-go:
	SRC_YAML=$(SRC_YAML) $(MAKE) -C $(CLIENT_PACKAGE) generate

##################
# Client targets #
##################
clean-client-js:
	$(MAKE) -C $(JS_CLIENT_PACKAGE) clean

generate-client-js:
	SRC_YAML=$(SRC_YAML) $(MAKE) -C $(JS_CLIENT_PACKAGE) generate

#################
# Model targets #
#################
clean-model:
	$(MAKE) -C $(MODEL_PACKAGE) clean

generate-model:
	SRC_YAML=$(SRC_YAML) $(MAKE) -C $(MODEL_PACKAGE) generate


############
# CLI tool #
############

build-cli:
	go build \
		-a \
		-ldflags "-X github.com/getlunaform/lunaform/cli/cmd.version=$(VERSION)" \
		-installsuffix $(CGO) \
		-o $(CWD)/lunaform \
		github.com/getlunaform/lunaform/cli


################
# Test targets #
################
test:
	go tool vet $(GO_TARGETS)
	go test $(GOR_TARGETS)

test-coverage:
	@sh $(CWD)/scripts/test-coverage.sh $(CWD) "$(GO_TARGETS)"
	go tool cover -html=$(CWD)/profile.out -o $(CWD)/coverage.html

validate-swagger:
	swagger validate $(SRC_YAML)

format:
	go fmt $(shell go list ./...)

lint:
	diff -u <(echo -n) <(gofmt -d -s $(shell find server -type d))
	diff -u <(echo -n) <(gofmt -d -s $(shell find backend -type d))
	golint -set_exit_status . $(GOR_TARGETS)


##################
# Docker targets #
##################
build-docker:
	GOOS=linux $(MAKE) lunaform
	docker build -t lunaform .

run-docker: build-docker
	docker run -p 8080:8080 lunaform