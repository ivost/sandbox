NAME=boss_pod_template
VERSION?=1.0.0
GIT_COMMIT:=`(test -d .git  && git describe --dirty --always)`
GIT_COMMIT?="unknown"
LDFLAGS=-ldflags "-X main.VERSION=${VERSION} -X main.REVISION=${GIT_COMMIT}"

define msg
	@printf "\033[36m $1 \n\033[0m"
endef

define msgc
	@printf "\n\033[32m\xE2\x9c\x93 $1 \n\033[0m"
	@printf "\n"
endef

.PHONY: help
help: ## - Show help
	$(call msgc,"usage: make [target]")
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

##########
# GO     #
##########

.PHONY: build
build: ## - Build app
	go build -o ${GOPATH}/bin/${NAME} ${LDFLAGS} cmd/*.go

.PHONY: run
run: ## - Run app
	go run cmd/*.go
