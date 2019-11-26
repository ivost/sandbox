SHELL := /bin/bash
BUILD_TARGET = build
GO := GO111MODULE=on go
GO_NOMOD :=GO111MODULE=off go
REV := $(shell git rev-parse --short HEAD 2> /dev/null || echo 'unknown')
ORG := ivostoy
ORG_REPO := $(ORG)/$(REPO)
RELEASE_ORG_REPO := $(ORG_REPO)
ROOT_PACKAGE := github.com/$(ORG_REPO)
GO_VERSION := $(shell $(GO) version | sed -e 's/^[^0-9.]*\([0-9.]*\).*/\1/')
GO_DEPENDENCIES := $(call rwildcard,pkg/,*.go) $(call rwildcard,cmd/,*.go)

BRANCH     := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown')
BUILD_DATE := $(shell date +%Y%m%d-%H:%M:%S)
CGO_ENABLED = 0

REPORTS_DIR=$(BUILD_TARGET)/reports

# Make does not offer a recursive wildcard function, so here's one:
rwildcard=$(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2))

GOTEST := $(GO) test
# If available, use gotestsum which provides more comprehensive output
# This is used in the CI builds
ifneq (, $(shell which gotestsum 2> /dev/null))
GOTESTSUM_FORMAT ?= standard-quiet
GOTEST := GO111MODULE=on gotestsum --junitfile $(REPORTS_DIR)/integration.junit.xml --format $(GOTESTSUM_FORMAT) --
endif

# requires git tag
VERSION ?= $(shell echo "$$(git for-each-ref refs/tags/ --count=1 --sort=-version:refname --format='%(refname:short)' 2>/dev/null)-dev+$(REV)")

BUILDFLAGS :=  -ldflags \
  " -X $(ROOT_PACKAGE)/pkg/version.Version=$(VERSION)\
		-X $(ROOT_PACKAGE)/pkg/version.Revision='$(REV)'\
		-X $(ROOT_PACKAGE)/pkg/version.Branch='$(BRANCH)'\
		-X $(ROOT_PACKAGE)/pkg/version.BuildDate='$(BUILD_DATE)'\
		-X $(ROOT_PACKAGE)/pkg/version.GoVersion='$(GO_VERSION)'"

ifdef DEBUG
BUILDFLAGS := -gcflags "all=-N -l" $(BUILDFLAGS)
endif

ifdef PARALLEL_BUILDS
BUILDFLAGS += -p $(PARALLEL_BUILDS)
GOTEST += -p $(PARALLEL_BUILDS)
else
# -p 4 seems to work well for people
GOTEST += -p 4
endif

ifdef DISABLE_TEST_CACHING
GOTEST += -count=1
endif

TEST_PACKAGE ?= ./...
COVER_OUT:=$(REPORTS_DIR)/cover.out
COVERFLAGS=-coverprofile=$(COVER_OUT) --covermode=count --coverpkg=./...

define msg
	@printf "\033[36m $1 \n\033[0m"
endef

define msgc
	@printf "\n\033[32m\xE2\x9c\x93 $1 \n\033[0m"
	@printf "\n"
endef

.PHONY: list
list: ## List all make targets
	@$(MAKE) -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$2 !~ "^[#.]") {print $$2}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: help
.DEFAULT_GOAL := help
help: ## Show help message
	$(call msgc,"usage: make [target]")
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: get-test-deps
get-test-deps: ## Install test dependencies
	$(GO_NOMOD) get github.com/axw/gocov/gocov
	$(GO_NOMOD) get -u gopkg.in/matm/v1/gocov-html

.PHONY: tidy-deps
tidy-deps: ## Cleans up dependencies
	$(GO) mod tidy
	@# mod tidy only takes compile dependencies into account, let's make sure we capture tooling dependencies as well
	@#@$(MAKE) install-generate-deps

.PHONY: make-reports-dir
make-reports-dir:
	mkdir -p $(REPORTS_DIR)

test: make-reports-dir ## Run ONLY the tests that have no build flags (and example of a build tag is the "integration" build tag)
	CGO_ENABLED=$(CGO_ENABLED) $(GOTEST) $(COVERFLAGS) -failfast -short ./...

test-report-html: make-reports-dir get-test-deps test ## Create the test report in HTML format
	@gocov convert $(COVER_OUT) | gocov-html > $(REPORTS_DIR)/cover.html && open $(REPORTS_DIR)/cover.html

test-verbose: make-reports-dir ## Run the unit tests in verbose mode
	CGO_ENABLED=$(CGO_ENABLED) $(GOTEST) -v $(COVERFLAGS) -failfast ./...

test-slow-report: get-test-deps test-slow make-reports-dir
	@gocov convert $(COVER_OUT) | gocov report

test-slow: make-reports-dir ## Run unit tests sequentially
	@CGO_ENABLED=$(CGO_ENABLED) $(GOTEST) $(COVERFLAGS) ./...

.PHONY: testr
testr: ## Clean test cache, test with race
	go clean -testcache ./...
	go mod download
	go mod verify
	go test --race ./... -v

.PHONY: clean
clean: ## Clean caches and app
	go clean -testcache ./...
	rm build/$(NAME)

.PHONY: release
release: ## make release with goreleaser
	@echo make sure to commit, tag and push (git cannot be in dirty state)
	@echo i.e.: git commit -a -m 'releasing...' && git tag -a v0.1.0 -m "First release" && git push origin v0.1.0
	@echo for dry run: goreleaser --snapshot --skip-publish
	rm -rf dist
	goreleaser

.PHONY: dr
dr: ## dry run with goreleaser
	rm -rf dist
	goreleaser --snapshot --skip-publish
