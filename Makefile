# Adapted from https://github.com/bufbuild/buf-example/blob/main/Makefile

SHELL := /usr/bin/env bash -o pipefail

# This controls the location of the cache.
PROJECT := galasa-api-go

OPENAPI_GENERATOR_VERSION := 7.8.0

GALASA_OPENAPI_VERSION := 0.37.0
GALASA_OPENAPI_YAML := openapi/openapi.yaml

# This is the commit hash for the https://github.com/googleapis/googleapis repo
# GRPC_STATUS_VERSION := f36c65081b19e0758ef5696feca27c7dcee5475e
# GRPC_STATUS_PROTO := google/rpc/status.proto

### Everything below this line is meant to be static, i.e. only adjust the above variables. ###

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)
ifeq ($(UNAME_OS),Darwin)
	PLATFORM := osx
	ifeq ($(UNAME_ARCH),arm64)
		PROTOC_ARCH := aarch_64
	else
		PROTOC_ARCH := x86_64
	endif
else
	PROTOC_ARCH := $(UNAME_ARCH)
endif
ifeq ($(UNAME_OS),Linux)
	PLATFORM := linux
endif
# Dependencies will be cached to ~/.cache/galasa-api-go.
CACHE_BASE := $(HOME)/.cache/$(PROJECT)
# This allows switching between i.e a Docker container and your local setup without overwriting.
CACHE := $(CACHE_BASE)/$(UNAME_OS)/$(UNAME_ARCH)
# The location where dependencies will be installed.
CACHE_BIN := $(CACHE)/bin
# Marker files are put into this directory to denote the current version of binaries that are installed.
CACHE_VERSIONS := $(CACHE)/versions

# Update the $PATH so we can use dependencies directly
export PATH := $(abspath $(CACHE_BIN)):$(PATH)
# Update GOBIN to point to CACHE_BIN for source installations
export GOBIN := $(abspath $(CACHE_BIN))
# This is needed to allow versions to be added to Golang modules with go get
export GO111MODULE := on

# OPENAPI_GENERATOR points to the marker file for the installed version.
#
# If OPENAPI_GENERATOR_VERSION is changed, the binary will be re-downloaded.
OPENAPI_GENERATOR := $(CACHE_VERSIONS)/openapi-generator-cli/$(OPENAPI_GENERATOR_VERSION)
$(OPENAPI_GENERATOR):
	@rm -f $(CACHE_BIN)/openapi-generator-cli
	@mkdir -p $(CACHE_BIN)
	curl -sSL \
		"https://raw.githubusercontent.com/OpenAPITools/openapi-generator/v$(OPENAPI_GENERATOR_VERSION)/bin/utils/openapi-generator-cli.sh" \
		-o "$(CACHE_BIN)/openapi-generator-cli"
	chmod +x "$(CACHE_BIN)/openapi-generator-cli"
	@rm -rf $(dir $(OPENAPI_GENERATOR))
	@mkdir -p $(dir $(OPENAPI_GENERATOR))
	@touch $(OPENAPI_GENERATOR)

.DEFAULT_GOAL := all

.PHONY: all
all: genapi

# deps allows us to install deps without running any checks.

.PHONY: deps
deps: $(OPENAPI_GENERATOR)

# Generate the Go client API for Galasa

$(GALASA_OPENAPI_YAML):
	@mkdir -p $(dir $(GALASA_OPENAPI_YAML))
	curl -sSL \
		"https://raw.githubusercontent.com/galasa-dev/framework/v$(GALASA_OPENAPI_VERSION)/galasa-parent/dev.galasa.framework.api.openapi/src/main/resources/openapi.yaml" \
		-o "$(GALASA_OPENAPI_YAML)"

.PHONY: genapi
genapi: deps $(GALASA_OPENAPI_YAML)
	OPENAPI_GENERATOR_VERSION=$(OPENAPI_GENERATOR_VERSION) openapi-generator-cli generate \
		-i ${GALASA_OPENAPI_YAML} \
		-g go \
		-o pkg/galasaapi \
		--additional-properties=packageName=galasaapi \
		--additional-properties=isGoSubmodule=false \
		--global-property apis,apiDocs=false,apiTests=false

# clean deletes any files not checked in and the cache for all platforms.

.PHONY: clean
clean:
	git clean -xdf

.PHONY: cleandep
cleandep:
	rm -rf $(CACHE_BASE)
