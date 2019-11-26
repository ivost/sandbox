NAME=boss-test-pod
REGION=us-west-2
DOCKER_REPO=930325689781.dkr.ecr.${REGION}.amazonaws.com
GIT_COMMIT=`git describe --dirty --always`
# VERSION:=`grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')`
VERSION?=1.0.0
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${GIT_COMMIT}"
TAG=${VERSION}_${GIT_COMMIT}
TAG?=latest
IMG=${DOCKER_REPO}/${NAME}:${TAG}
# BUILD=`date +%FT%T%z`

all: ## - all
	@echo all
	$(call msg,"all")
	$(call msgc,"all")


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
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ---
	@echo GIT_COMMIT: ${GIT_COMMIT}
	@echo IMG: ${IMG}
	@echo ---

##########
# Docker #
##########

.PHONY: dbuild
dbuild:	## - Build docker image
	$(call msgc,"Building docker image...")
	@docker build --file Dockerfile --tag ${IMG} .

.PHONY: dbuild-no-cache
dbuild-no-cache: ## - docker build without cache
	$(call msgc,"Building docker image without cache...")
	@export DOCKER_CONTENT_TRUST=1 && docker build --no-cache -f Dockerfile --tag  $(IMG) .

.PHONY: dls
dls: ## - List image size
	$(call msgc,"Show docker image size $(IMG)")
	@docker image ls $(IMG)

.PHONY: drun
drun: ## - Run docker container
	$(call msgc,"Run docker container")
	@docker run  $(IMG)
	@docker run --rm -d $(IMG) -p 8080:8080

.PHONY: dpush
dpush: ## - Push docker container to ECR
	@`aws ecr get-login --region ${REGION} --no-include-email`
	docker push $(IMG)

# docker tag boss-go-builder:latest 930325689781.dkr.ecr.us-west-2.amazonaws.com/boss-go-builder:latest
# docker tag d7fe1a9d75d5 930325689781.dkr.ecr.us-west-2.amazonaws.com/boss-go-builder:latest
# docker push 930325689781.dkr.ecr.us-west-2.amazonaws.com/boss-go-builder:latest
# docker tag d7fe1a9d75d5 930325689781.dkr.ecr.us-west-2.amazonaws.com/boss-go-builder:1.12
# docker push 930325689781.dkr.ecr.us-west-2.amazonaws.com/boss-go-builder:1.12
