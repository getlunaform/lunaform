.DEFAULT_GOAL := default

SRC_YAML?="swagger.yml"
CGO?=cgo

CWD?=$(shell pwd)

SHELL:=/bin/bash
ENVIRONMENT?=DEVELOPMENT

BUILD_ID?=$(ENVIRONMENT)

GOROOT?=${HOME}/go

GO_TARGETS= ./server ./backend
GOR_TARGETS= ./server/... ./backend/...

VERSION?=$(shell git rev-parse --short HEAD)

MODEL_PATH?=${GOROOT}/src/github.com/getlunaform/lunaform-models-go

##################
# Global Targets #
##################
build: build-server build-client
clean: clean-server clean-client
generate: generate-server generate-client

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
		github.com/drewsonne/lunaform/server/cmd/lunaform-server

clean-server:
	cp $(MODEL_PATH)/hal.go $(CWD)/hal.go && \
	rm -rf $(CWD)/server/cmd/ \
		$(CWD)/server/restapi/operations \
		$(CWD)/server/restapi/doc.go \
		$(CWD)/server/restapi/embedded_spec.go \
		$(CWD)/server/restapi/server.go \
		$(CWD)/lunaform \
		$(CWD)/profile.txt && \
	mv $(CWD)/hal.go $(MODEL_PATH)/hal.go

generate-server:
	swagger generate server \
		--target=server \
		--principal=models.Principal \
		--name=lunaform \
		--skip-models \
		--existing-models github.com/getlunaform/lunaform-models-go \
		--spec=$(SRC_YAML)

run-server:
	$(CWD)/lunaform --port=8080 --scheme=http

##################
# Client targets #
##################

build-client:
	go build \
		-ldflags "-X github.com/drewsonne/lunaform/cli/cmd.version=$(VERSION)" \
		-a -installsuffix $(CGO) \
		-o $(CWD)/lunaform \
		github.com/drewsonne/lunaform/cli

clean-client:
	rm -f $(CWD)/lunaform && \
	rm -rf $(CWD)/client && \
	mkdir -p $(CWD)/client && \
	touch $(CWD)/client/.gitkeep

generate-client:
	swagger generate client \
		-f swagger.yml \
		-A lunaform \
		--client-package=lunaform-client-go \
		--target=${GOROOT}/src/github.com/getlunaform/ \
		--existing-models github.com/getlunaform/lunaform-models-go \
		--skip-models

#################
# Model targets #
#################

generate-models:
	swagger generate model \
		-f swagger.yml \
		--model-package=lunaform-models-go \
		--target=${GOROOT}/src/github.com/getlunaform/


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