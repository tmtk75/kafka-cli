.DEFAULT_GOAL := help

#VERSION := $(shell git describe --tags --abbrev=0)
#VERSION_LONG := $(shell git describe --tags)
VERSION := dev
VERSION_LONG := dev
#VAR_VERSION := github.com/tmtk75/kafka-cli/cmd.Version
VAR_VERSION := main.Version

LDFLAGS := -ldflags "-X $(VAR_VERSION)=$(VERSION) \
	-X $(VAR_VERSION)Long=$(VERSION_LONG)"

SRCS := $(shell find . -type f -name '*.go')

.PHONY: build
build: kafka-cli ## Build here

kafka-cli: $(SRCS)
	go build $(LDFLAGS) .

.PHONY: install
install:  ## Install in GOPATH
	go install $(LDFLAGS) .

.PHONY: clean
clean:  ## Clean
	rm -f kafka-cli

distclean: clean
	rm -rf build

## Release targets
.PHONY: build-release archive
build-release: build/kafka-cli_linux_amd64 build/kafka-cli_darwin_amd64
archive: build/kafka-cli_linux_amd64.gz build/kafka-cli_darwin_amd64.gz
release: upload-archives

upload-archives: archive
	ghr -u tmtk75 $(VERSION) ./build/*.gz

build/kafka-cli_linux_amd64.gz: build/kafka-cli_linux_amd64
	(cd build; gzip --keep kafka-cli_linux_amd64)

build/kafka-cli_darwin_amd64.gz: build/kafka-cli_darwin_amd64
	(cd build; gzip --keep kafka-cli_darwin_amd64)

build/kafka-cli_linux_amd64:
	GOARCH=amd64 GOOS=linux go build $(LDFLAGS) -o build/kafka-cli_linux_amd64 .

build/kafka-cli_darwin_amd64:
	GOARCH=amd64 GOOS=darwin go build $(LDFLAGS) -o build/kafka-cli_darwin_amd64 .

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

