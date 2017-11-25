update-vendor:
	glide update

format:
	go fmt $(shell go list ./...)

lint:
	golint $(shell go list ./... | grep -v vendor)
