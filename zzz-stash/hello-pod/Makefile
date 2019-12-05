NAME=hello
SHELL := /bin/bash
BUILD_TARGET = build
GO := GO111MODULE=on go
GO_NOMOD :=GO111MODULE=off go

#REGION=us-west-2
#DOCKER_REPO=930325689781.dkr.ecr.${REGION}.amazonaws.com


# dockerhub
DOCKER_REPO=ivostoy
DAEMON=
VERSION=0.11.29.0
#REV := $(shell git rev-parse --short HEAD 2> /dev/null || echo 'unknown')
#ORG := ivostoy
#ORG_REPO := $(ORG)/$(REPO)

IMG=${DOCKER_REPO}/${NAME}:${TAG}

# Make does not offer a recursive wildcard function, so here's one:
rwildcard=$(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2))

#RELEASE_ORG_REPO := $(ORG_REPO)
#ROOT_PACKAGE := github.com/$(ORG_REPO)
GO_VERSION := $(shell $(GO) version | sed -e 's/^[^0-9.]*\([0-9.]*\).*/\1/')
GO_DEPENDENCIES := $(call rwildcard,pkg/,*.go) $(call rwildcard,cmd/,*.go)

BRANCH     := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown')
BUILD_DATE := $(shell date +%Y%m%d-%H:%M:%S)
GIT_COMMIT=$(shell git describe --dirty --always  2> /dev/null  || echo 'unknown')
BUILDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${GIT_COMMIT}"
CGO_ENABLED = 0
MAIN_SRC_FILE=main.go
# to check - used by kustomize
#TAG=${VERSION}_${GIT_COMMIT}
IMG_TAG=${VERSION}
IM=${DOCKER_REPO}/${NAME}
IMG=${DOCKER_REPO}/${NAME}:${IMG_TAG}
BASE=./kustomize/base
OVERLAYS=./kustomize/overlays

#include common.mk

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



.PHONY: build
build: $(GO_DEPENDENCIES) ## Build binary for the current OS
	CGO_ENABLED=$(CGO_ENABLED) $(GO) $(BUILD_TARGET) $(BUILDFLAGS) -o $(NAME) $(MAIN_SRC_FILE)

##########
# Docker #
##########

.PHONY: dbuild
dbuild:	## - Build docker image
	$(call msgc,"Building docker image...")
	@docker build --file Dockerfile --tag ${IMG} .

.PHONY: dbuild-no-cache
dbuild-no-cache: ## - Docker build without cache
	$(call msgc,"Building docker image without cache...")
	@export DOCKER_CONTENT_TRUST=1 && docker build --no-cache -f Dockerfile --tag  $(IMG) .

.PHONY: dsize
dsize: ## - Show docker image size
	$(call msgc,"Show docker image size")
	@docker image ls $(IMG)

.PHONY: drun
drun: ## - Run docker container
	$(call msgc,"Run docker container")
	@echo when running in minikube - check you DOCKER_HOST minikube docker-env
	@echo i.e. http://192.168.99.100:8080
	@echo ---
	@docker run --rm $(DAEMON) -p 8080:8080 --name $(NAME) $(IMG)

.PHONY: dkill
dkill: ## - Kill running docker container
	$(call msgc,"Kill running docker container")
	@docker kill $(NAME)

.PHONY: dps
dps: ## - Container status
	docker ps | grep $(NAME)

.PHONY: dpush
dpush: ## - Publish image $(IMG)
	## @`aws ecr get-login --region ${REGION} --no-include-email`
	docker push $(IMG)

.PHONY: dpull
dpull: ## - Pull docker container from ECR
	## @`aws ecr get-login --region ${REGION} --no-include-email`
	docker pull $(IMG)


##############
# Kubernetes #
##############

.PHONY: kdeploy
kdeploy: ## Kustomize and Deploy to k8s (test2)
	cd $(BASE) && kustomize edit set image $(IM)=$(IMG) | kustomize build . | kubectl apply -f - && cd ..
	#kustomize build $(OVERLAYS)/test2
	#kustomize build $(OVERLAYS)/test2 | kubectl apply -f -
	sleep 1

.PHONY: kundeploy
kundeploy: ## Undeploy from k8s
	# disable delete errors
	@set +e
	cd $(BASE) && kubectl delete -f deployment.yaml && cd ..
	#kustomize build $(OVERLAYS)/test2 | kubectl delete -f -
	@set -e
	sleep 1


