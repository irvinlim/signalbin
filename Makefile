SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: clean build

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

SIGNALBIN ?= bin/signalbin

.PHONY: clean
clean: ## Clean all built targets.
	rm -f $(SIGNALBIN)

.PHONY: build
build: $(SIGNALBIN) ## Build signalbin.
$(SIGNALBIN):
	go build -o $(SIGNALBIN) ./...
