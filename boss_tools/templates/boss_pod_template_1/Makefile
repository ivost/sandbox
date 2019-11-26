NAME=boss_pod_template
REGION=us-west-2
DOCKER_REPO=930325689781.dkr.ecr.${REGION}.amazonaws.com
GIT_COMMIT=`git describe --dirty --always`
# VERSION:=`grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')`
VERSION?=1.0.0
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${GIT_COMMIT}"
TAG=${VERSION}_${GIT_COMMIT}
TAG?=latest
####
TAG=latest
####
IMG=${DOCKER_REPO}/${NAME}:${TAG}
# BUILD=`date +%FT%T%z`

#DAEMON=-d
DAEMON=
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
	@echo ---
	@echo IMAGE: ${IMG}
	@echo ---


.PHONY: archive
archive: ## - Create zip archive
	mkdir -p archive
	zip -x *.idea* -x *.git* *archive*  @  -r archive/master .

##########
# GO     #
##########

.PHONY: test
test: ## - Test app
	go clean -testcache ./...
	go test --race ./... -v

.PHONY: build
build: ## - Build app
	go build -o ${GOPATH}/bin/podinfo cmd/podinfo/main.go
	go build -o ${GOPATH}/bin/podcli cmd/podcli/*

.PHONY: run
run: ## - Run app
	go run cmd/podinfo/main.go

.PHONY: docs
docs: ## - Generate swagger doc
	go get github.com/swaggo/swag/cmd/swag
	swag init -g ./pkg/api/docs/docs.go

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
dpush: ## - Push docker container to ECR
	@`aws ecr get-login --region ${REGION} --no-include-email`
	docker push $(IMG)

.PHONY: dpull
dpull: ## - Pull docker container from ECR
	@`aws ecr get-login --region ${REGION} --no-include-email`
	docker pull $(IMG)

# docker tag boss-go-builder:latest 930325689781.dkr.ecr.us-west-2.amazonaws.com/boss-go-builder:latest
# docker tag d7fe1a9d75d5 930325689781.dkr.ecr.us-west-2.amazonaws.com/boss-go-builder:latest
# docker push 930325689781.dkr.ecr.us-west-2.amazonaws.com/boss-go-builder:latest

# swag init -g ./pkg/api/docs/docs.go