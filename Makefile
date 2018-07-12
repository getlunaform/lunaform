SRC_YAML?="swagger.yml"
CGO?=cgo

CWD?=$(shell pwd)

SHELL:=/bin/bash
GO_PIPELINE_LABEL?=BUILD_ID
ENVIRONMENT?=DEVELOPMENT

BUILD_NUMBER?=$(GO_PIPELINE_LABEL)
BUILD_ID?=$(ENVIRONMENT)

GO_TARGETS= ./server ./backend
GOR_TARGETS= ./server/... ./backend/...

SHA?=$(shell git rev-parse HEAD)
BUILT_BY?=$(shell whoami)
HOSTNAME?=$(shell hostname)
NOW?=$(shell date +%s)

VERSION?=$(shell git rev-parse --short HEAD)

build: generate lunaform

run-clean: clean build run

run: lunaform
	$(CWD)/lunaform --port=8080 --scheme=http

update-vendor:
	glide update

clean: clean-client
	cp $(CWD)/server/models/hal.go $(CWD)/hal.go && \
	rm -rf $(CWD)/server/cmd/ \
		$(CWD)/server/models/ \
		$(CWD)/server/restapi/operations \
		$(CWD)/server/restapi/doc.go \
		$(CWD)/server/restapi/embedded_spec.go \
		$(CWD)/server/restapi/server.go \
		$(CWD)/lunaform \
		$(CWD)/profile.txt && \
	mkdir -p $(CWD)/server/models && \
	mv $(CWD)/hal.go $(CWD)/server/models/hal.go


clean-client:
	rm -f $(CWD)/lunaform-client && \
	rm -rf $(CWD)/client

validate-swagger:
	swagger validate $(SRC_YAML)


test:
	go tool vet $(GO_TARGETS)
	go test $(GOR_TARGETS)

test-coverage:
	@sh $(CWD)/scripts/test-coverage.sh $(CWD) "$(GO_TARGETS)"
	go tool cover -html=$(CWD)/profile.out -o $(CWD)/coverage.html

format:
	go fmt $(shell go list ./...)

lint:
	diff -u <(echo -n) <(gofmt -d -s $(shell find server -type d))
	diff -u <(echo -n) <(gofmt -d -s $(shell find backend -type d))
	golint -set_exit_status . $(GOR_TARGETS)

generate-swagger: validate-swagger
	swagger generate server \
		--target=server \
		--principal=models.Principal \
		--name=lunaform \
		--spec=$(SRC_YAML)

generate: generate-swagger generate-client

lunaform:
	go build \
		-a -installsuffix $(CGO) \
		-o ./lunaform \
		github.com/drewsonne/lunaform/server/cmd/lunaform-server

build-docker:
	GOOS=linux $(MAKE) lunaform
	docker build -t lunaform .

run-docker: build-docker
	docker run -p 8080:8080 lunaform

generate-client:
	mkdir -p client && \
	swagger generate client \
		-f swagger.yml \
		-A lunaform-client \
		--existing-models github.com/drewsonne/lunaform/server/models \
		--skip-models \
		--target client

build-client: generate-client
	go build -ldflags "-X github.com/drewsonne/lunaform/cli/cmd.version=$(VERSION)" -o lunaform-client github.com/drewsonne/lunaform/cli

client-clean: clean-client build-client