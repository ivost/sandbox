BUILD_TARGET = build
#GO := GO111MODULE=on GOOS=linux GOARCH=amd64 go
GO := go
GO_NOMOD := GO111MODULE=off go
NAME := myservice

# dockerhub
DOCKER_REPO=ivostoy
DAEMON=
VERSION=0.12.3.0

# Make does not offer a recursive wildcard function, so here's one:
rwildcard=$(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2))

#RELEASE_ORG_REPO := $(ORG_REPO)
#ROOT_PACKAGE := github.com/$(ORG_REPO)
GO_VERSION := $(shell $(GO) version | sed -e 's/^[^0-9.]*\([0-9.]*\).*/\1/')
GO_DEPENDENCIES := $(call rwildcard,pkg/,*.go) $(call rwildcard,cmd/,*.go)

BRANCH     := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown')
BUILD_DATE := $(shell date +%Y%m%d-%H:%M:%S)
PKG=github.com/ivost/sandbox/myservice/pkg
GIT_COMMIT=$(shell git describe --dirty --always  2> /dev/null  || echo 'unknown')
BUILDFLAGS=-ldflags "-s -w -X ${PKG}/version.Version=${VERSION} -X ${PKG}/version.Build=${GIT_COMMIT}"
CGO_ENABLED = 0
MAIN_SRC_FILE=main.go
# to check - used by kustomize
#TAG=${VERSION}_${GIT_COMMIT}
IMG_TAG=${VERSION}
IM=${DOCKER_REPO}/${NAME}
IMG=${DOCKER_REPO}/${NAME}:${IMG_TAG}

IMG_BLD_TAG=build_${VERSION}
IMG_BLD=${DOCKER_REPO}/${NAME}:${IMG_BLD_TAG}

BASE=./kustomize/base
OVERLAYS=./kustomize/overlays
POD:=$(shell kubectl get pod -l app=hello -o  jsonpath='{.items[*].metadata.name}') > /dev/nul

# kubectl get pods -l app=hello -o=jsonpath="{range .items[*]}{.metadata.name}{"\n"}{end}"

# kubectl get pods -l app=hello -o jsonpath="{.items[*]['metadata.name']}"

# kubectl get pods -l app=ambassador -o jsonpath="{.items[*]['metadata.name']}"

#include common.mk

define msg
	@printf "\033[36m $1 \n\033[0m"
endef

define msgc
	@printf "\n\033[32m\xE2\x9c\x93 $1 \n\033[0m"
	@printf "\n"
endef

#.PHONY: all
.DEFAULT_GOAL := all
all: dbuild dpush kdeploy		## docker dbuild / dpush / kdeploy

#all: dbuild dpush kundeploy kdeploy		## docker dbuild / dpush / kdeploy

.PHONY: list
list: ## List all make targets
	@$(MAKE) -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$2 !~ "^[#.]") {print $$2}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: help
### .DEFAULT_GOAL := help
help: ## Show help message
	$(call msgc,"usage: make [target]")
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


.PHONY: run
run: build		## Build and run server (local)
	build/server

.PHONY: build
build:  ## Build binary for the current OS
	./proto-gen.sh
	#go env
	go $(BUILD_TARGET) $(BUILDFLAGS) -o build/server cmd/server/server.go
	go $(BUILD_TARGET) $(BUILDFLAGS) -o build/client cmd/client/client.go

.PHONY: build.linux
build-linux:  ## Build binary for linux (2phase build is very slow)
	GOOS=linux go $(BUILD_TARGET) $(BUILDFLAGS) -o lbuild/server cmd/server/server.go
	#GOOS=linux go $(BUILD_TARGET) $(BUILDFLAGS) -o lbuild/client cmd/client/client.go
	#chmod +x lbuild/server
	#chmod +x lbuild/client

##########
# Docker #
##########

.PHONY: dbuild
dbuild:	build-linux ## - Build local - copy to docker image
	$(call msgc,"Building docker image $(IMG) ...")
	docker build --tag $(IMG) .
	#docker push $(IMG)

.PHONY: dbuild.multi
dbuild.multi:	## - Build docker image in CI
	$(call msgc,"Building docker image...")
	#@docker build --file Dockerfile --tag ${IMG} .

	# Pull the latest version of the image, in order to
	# populate the build cache:
	docker pull $(IMG_BLD) || true
	docker pull $(IMG)     || true

	# Build the compile stage:
	docker build --target builder \
		   --cache-from=$(IMG_BLD) \
		   --tag $(IMG_BLD) .

	# Build the runtime stage, using cached compile stage:
	docker build --target runtime \
		   --cache-from=$(IMG_BLD) \
		   --cache-from=$(IMG) \
		   --tag $(IMG) .
	# Push the new versions:
	docker push $(IMG_BLD)
	docker push $(IMG)

.PHONY: dbuild-no-cache
dbuild-no-cache: ## - Docker build without cache
	$(call msgc,"Building docker image without cache...")
	@export DOCKER_CONTENT_TRUST=1 && docker build --no-cache -f Dockerfile --tag  $(IMG) .

.PHONY: dsize
dsize: ## - Show docker image size
	$(call msgc,"Show docker image size")
	@docker image ls $(IMG)

.PHONY: drun
drun: dbuild ## - Run docker container
	$(call msgc,"Run docker container")
	@echo when running in minikube - check you DOCKER_HOST minikube docker-env
	@echo i.e. http://192.168.99.100:8080
	@echo ---
	@docker run --rm $(DAEMON) --name $(NAME) -p 8080:8080 -p 52052:52052 $(IMG)

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

.PHONY: kpf
kpf: ## port forwarding
	kubectl port-forward $(POD) 52052:52052 8080:8080

.PHONY: kcall
kcall: ## call into k8s via ambassador
	build/client -config=client-config.yaml
