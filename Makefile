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

run-clean: clean build run

run: lunarform
	$(CWD)/lunarform --port=8080 --scheme=http

update-vendor:
	glide update

clean: clean-client
	rm -rf $(CWD)/server/cmd/ \
		$(CWD)/server/models/ \
		$(CWD)/server/restapi/operations \
		$(CWD)/server/restapi/doc.go \
		$(CWD)/server/restapi/embedded_spec.go \
		$(CWD)/server/restapi/server.go \
		$(CWD)/lunarform \
		$(CWD)/profile.txt \


clean-client:
	rm -f $(CWD)/tfs-client && \
	rm -rf $(CWD)/client

validate-swagger:
	swagger validate $(SRC_YAML)

build: generate-swagger lunarform

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
		--name=TerraformServer \
		--spec=$(SRC_YAML)

generate:generate-swagger

lunarform:
	go build \
		-a -installsuffix $(CGO) \
		-o ./lunarform \
		github.com/drewsonne/lunarform/server/cmd/lunarform-server

build-docker:
	GOOS=linux $(MAKE) lunarform
	docker build -t lunarform .

run-docker: build-docker
	docker run -p 8080:8080 lunarform

build-client:
	mkdir client && \
	swagger generate client \
		-f swagger.yml \
		-A lunarform-client \
		--existing-models github.com/drewsonne/lunarform/server/models \
		--skip-models \
		--target client && \
	go build -ldflags "-X github.com/drewsonne/lunarform/cli/cmd.version=$(VERSION)" -o tfs-client github.com/drewsonne/lunarform/cli

client-clean: clean-client build-client