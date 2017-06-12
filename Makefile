PACKAGES = $(shell go list ./... | grep -v /vendor/)
all: build
vet:
	go vet $(PACKAGES)
build:
	CGO_ENABLED=0 go build -o location.bin