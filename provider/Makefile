GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif

.PHONY: build
build: clean
	CGO_ENABLED=0 go build -ldflags "-s -w" -o bin .

.PHONY: run
run: build
	./bin

.PHONY: clean
clean:
	rm -f bin
	rm -rf polaris

.PHONY: build-docker
build-docker: build
	docker build -t polaris-go-bin:latest .
