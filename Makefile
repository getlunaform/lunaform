SRC_YAML?="swagger.yml"
CGO?="cgo"

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

LDFLAGS?="-X github.com/zeebox/terraform-server/server/restapi.builtWhen=$(NOW) \
			-X github.com/zeebox/terraform-server/server/restapi.buildMachine=$(HOSTNAME) \
			-X github.com/zeebox/terraform-server/server/restapi.buildNumber=$(BUILD_NUMBER) \
			-X github.com/zeebox/terraform-server/server/restapi.builtBy=$(BUILT_BY) \
			-X github.com/zeebox/terraform-server/server/restapi.compiler=$(CGO) \
			-X github.com/zeebox/terraform-server/server/restapi.sha=$(SHA)"

doc:
	@sh scripts/generate-doc.sh

update-vendor:
	glide update

run: terraform-server
	$(PWD)/terraform-server --scheme=http

validate-swagger:
	swagger validate $(SRC_YAML)

build: generate-swagger terraform-server

test:
	go tool vet $(GO_TARGETS)
	go test $(shell go list $(GOR_TARGETS))

test-coverage:
	goverage -v -race -coverprofile=profile.txt -covermode=atomic $(shell go list $(GOR_TARGETS))
	go tool cover -html=profile.txt -o coverage.html

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

terraform-server:
	go build \
		-a -installsuffix $(CGO) \
		-o ./terraform-server \
		github.com/zeebox/terraform-server/server/cmd/terraform-server-server

build-docker:
	GOOS=linux $(MAKE) terraform-server
	docker build -t terraform-server .

run-docker: build-docker
	docker run -p 8080:8080 terraform-server
