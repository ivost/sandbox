NAME=cli
REGION=us-west-2

define msg
	@printf "\033[36m $1 \n\033[0m"
endef

define msgc
	@printf "\n\033[32m\xE2\x9c\x93 $1 \n\033[0m"
	@printf "\n"
endef

.PHONY: help
help: ## - Show help message
	$(call msgc,"usage: make [target]")
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

##########
# GO     #
##########

.PHONY: test
test: ## - Test app
	#go clean -testcache ./...
	go test --race ./... -v

.PHONY: build
build: ## - Build app
	go build -o ${GOPATH}/bin/cli cmd/*

.PHONY: run
run: ## - Run app
	go run cmd/main.go

