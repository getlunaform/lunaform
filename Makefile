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

EXISTING_MODELS?=github.com/getlunaform/lunaform-models-go

##################
# Global Targets #
##################
build: build-server build-cli
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
		github.com/getlunaform/lunaform/cmd/lunaform-server

clean-server:
	rm -rf ${CURDIR}/cmd/ \
		${CURDIR}/restapi/operations \
		${CURDIR}/restapi/doc.go \
		${CURDIR}/restapi/embedded_spec.go \
		${CURDIR}/restapi/server.go \
		${CURDIR}/lunaform \
		${CURDIR}/profile.txt

generate-server:
	swagger generate server \
		-f $(SRC_YAML) \
		--name=lunaform \
		--principal=models.ResourceAuthUser \
		--skip-models

run-server:
	$(CWD)/lunaform --port=8080 --scheme=http

##################
# Client targets #
##################
clean-client-go:
	find ${CURDIR}/client -name "*.go" -delete && \
	find ${CURDIR}/client -name ".DS_STORE" -delete && \
	find ${CURDIR}/client/ -mindepth 1 -type d -empty -delete

generate-client-go:
	swagger generate client \
		-f $(SRC_YAML) \
		--name=lunaform \
		--principal=models.ResourceAuthUser \
		--default-produces=application/vnd.lunaform.v1+json

#################
# Model targets #
#################
clean-model:
	find ${CURDIR}/models -type f \( -name "*.go" -not -name hal.go \) -delete && \
	find ${CURDIR}/models -name ".DS_STORE" -delete && \
	find ${CURDIR}/models -type d -empty -delete

generate-model:
	swagger generate model \
		-f $(SRC_YAML)

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

##################
# Client targets #
##################
clean-client-js:
	$(MAKE) -C $(JS_CLIENT_PACKAGE) clean

generate-client-js:
	SRC_YAML=$(SRC_YAML) $(MAKE) -C $(JS_CLIENT_PACKAGE) generate

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
#	diff -u <(echo -n) <(gofmt -d -s $(shell find backend -type d))
	golint -set_exit_status . $(shell glide novendor)


##################
# Docker targets #
##################
build-docker:
	GOOS=linux $(MAKE) lunaform
	docker build -t lunaform .

run-docker: build-docker
	docker run -p 8080:8080 lunaform