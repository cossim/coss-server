PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

BUILDPLATFORM := linux/amd64,linux/arm64,linux/arm/v8

MAIN_FILE=cmd/main.go
NAME= ""
DIR := $(shell pwd)
IMG ?= hub.hitosea.com/cossim/${NAME}-${ACTION}:latest

BUILD_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
#BUILD_BRANCH := "main"
BUILD_COMMIT := ${shell git rev-parse HEAD}
#BUILD_COMMIT := "22193944514397212fe6a25189906ab9de49164a"
BUILD_TIME := ${shell date '+%Y-%m-%d %H:%M:%S'}
BUILD_GO_VERSION := $(shell go version | grep -o  'go[0-9].[0-9].*')
VERSION_PATH := "github.com/cossim/coss-server/pkg/version"
BUILD_PATH := ""
DOCKER_BUILD_PATH := ""
ACTION := ""

CONFIG_PATH := "deploy/docker/config"
DOCKER_COMPOSE_PATH := "deploy/docker"

INTERFACE_LIST ?= group msg relation storage user live

CONSUL_HOST := "127.0.0.1:8500"
CONSUL_SSL := false
CONSUL_TOKEN := ""

# 根据传入的 ACTION 参数设置 BUILD_PATH
ifeq ($(ACTION), interface)
	BUILD_PATH := ${DIR}/interface/${NAME}
	DOCKER_BUILD_PATH :="interface/${NAME}"
else ifeq ($(ACTION), service)
	BUILD_PATH := ${DIR}/service/${NAME}
	DOCKER_BUILD_PATH := "service/${NAME}"
else ifeq ($(ACTION), admin)
	BUILD_PATH := ${DIR}/admin
	DOCKER_BUILD_PATH := "${NAME}"
endif

# 如果没有设置 BUILD_PATH，输出错误信息
ifeq ($(BUILD_PATH),)
    $(error Invalid ACTION. Use 'make build ACTION=interface' or 'make build ACTION=service')
endif

.PHONY: build-service build-interface docker-build docker-push

.PHONY: dep
dep: ## Get the dependencies
	@go mod tidy

.PHONY: lint
lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: test
test: fmt vet## Run unittests
	@go test -short ./...

SERVICE_DIR := ./service
INTERFACE_DIR := ./interface

.PHONY: run run_service run_interface

.PHONY: run stop

run: config_init
	@docker-compose -f $(DOCKER_COMPOSE_PATH)/docker-compose.yaml up -d

restart:
	@docker-compose -f $(DOCKER_COMPOSE_PATH)/docker-compose.yaml restart
stop:
	@docker-compose -f $(DOCKER_COMPOSE_PATH)/docker-compose.yaml down




#run_service:
#	@for dir in $(shell ls -d $(SERVICE_DIR)/*); do \
#		if [ -f "$$dir/Makefile" ]; then \
#			(cd $$dir && make run &); \
#		else \
#			echo "No Makefile found in $$dir"; \
#		fi \
#	done
#
#run_interface:
#	@for dir in $(shell ls -d $(INTERFACE_DIR)/*); do \
#		if [ -f "$$dir/Makefile" ]; then \
#			(cd $$dir && make run &); \
#		else \
#			echo "No Makefile found in $$dir"; \
#		fi \
#	done

swag: ## Run unittests
	$(foreach dir,$(INTERFACE_LIST), \
		swag i -g http.go -dir interface/$(dir)/server/http,interface/$(dir)/api/model,interface/live/api/dto,pkg/utils/usersorter --instanceName $(dir); \
	)
	swag i -g http.go -dir admin/server/http,admin/api/model,pkg/utils/usersorter --instanceName admin


.PHONY: config_init config_clear
config_init: ## Initialize config files
	@config_path=$(CONFIG_PATH) && \
	for template_file in $$(find $${config_path} -type f -name '*.template'); do \
		if [ -e "$$template_file" ]; then \
			new_file="$${template_file%.template}"; \
			if [ ! -e "$$new_file" ]; then \
				cp "$$template_file" "$$new_file"; \
				echo "Copied $$template_file to $$new_file"; \
			else \
				echo "File $$new_file already exists, skipping copy"; \
			fi \
		fi \
	done

config_clear: ## Clear YAML and JSON files in config_path
	@config_path=$(CONFIG_PATH) && \
	find $${config_path} \( -name '*.yaml' -o -name '*.yml' -o -name '*.json' \) -type f -exec rm -f {} \; && \
	echo "Cleared YAML and JSON files in $${config_path}"

#.PHONY: common dev prod
#common: config_init
#	@docker-compose -f $(DOCKER_COMPOSE_PATH)/docker-compose.yaml.bak2 up -d
#
#dev: common
#	@docker-compose -f $(DOCKER_COMPOSE_PATH)/docker-compose.yaml.bak2 -f $(DOCKER_COMPOSE_PATH)/docker-compose.dev.yaml up -d
#
#prod: common
#	@docker-compose -f $(DOCKER_COMPOSE_PATH)/docker-compose.yaml.bak2 -f $(DOCKER_COMPOSE_PATH)/docker-compose.prod.yaml up -d
#
#stop:
#	@docker-compose -f $(DOCKER_COMPOSE_PATH)/docker-compose.yaml.bak2 -f $(DOCKER_COMPOSE_PATH)/docker-compose.yaml.bak2 -f $(DOCKER_COMPOSE_PATH)/docker-compose.prod.yaml down

# 构建指定grpc服务  make build-service ACTION=service NAME="user"
build-service: dep ## Build the binary file
ifdef NAME
	@echo "Building with flags: go build -ldflags \"-s -w\" -ldflags \"-X '${VERSION_PATH}.GitBranch=${BUILD_BRANCH}' -X '${VERSION_PATH}.GitCommit=${BUILD_COMMIT}' -X '${VERSION_PATH}.BuildTime=${BUILD_TIME}' -X '${VERSION_PATH}.GoVersion=${BUILD_GO_VERSION}'\" -o ${BUILD_PATH}/$(MAIN_FILE)"
	@go build -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GitBranch=${BUILD_BRANCH}' -X '${VERSION_PATH}.GitCommit=${BUILD_COMMIT}' -X '${VERSION_PATH}.BuildTime=${BUILD_TIME}' -X '${VERSION_PATH}.GoVersion=${BUILD_GO_VERSION}'" -o ${BUILD_PATH}/bin/main ${BUILD_PATH}/$(MAIN_FILE)
else
	@echo "Please provide service NAME"
endif

# 构建指定接口服务  make build-interface ACTION=interface NAME="user"
build-interface: dep
ifdef NAME
	@echo "Building ${INTERFACE_NAME} interface with flags: go build -ldflags \"-s -w\" -ldflags \"-X '${VERSION_PATH}.GitBranch=${BUILD_BRANCH}' -X '${VERSION_PATH}.GitCommit=${BUILD_COMMIT}' -X '${VERSION_PATH}.BuildTime=${BUILD_TIME}' -X '${VERSION_PATH}.GoVersion=${BUILD_GO_VERSION}'\" -o ${BUILD_PATH}/$(MAIN_FILE)"
	@go build -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GitBranch=${BUILD_BRANCH}' -X '${VERSION_PATH}.GitCommit=${BUILD_COMMIT}' -X '${VERSION_PATH}.BuildTime=${BUILD_TIME}' -X '${VERSION_PATH}.GoVersion=${BUILD_GO_VERSION}'" -o ${BUILD_PATH}/bin/main ${BUILD_PATH}/$(MAIN_FILE)
else
	@echo "Please provide interface NAME"
endif

# If you wish built the manager image targeting other platforms you can use the --platform flag.
# (i.e. docker build --platform linux/arm64 ). However, you must enable docker buildKit for it.
# More info: https://docs.docker.com/develop/develop-images/build_enhancements/
docker-build: dep test## Build docker image with the manager.
	#docker build -t ${IMG} .
	# 根据传入的 ACTION 参数设置 BUILD_PATH
	docker build --build-arg BUILD_BRANCH="${BUILD_BRANCH}" \
             --build-arg BUILD_COMMIT="${BUILD_COMMIT}" \
             --build-arg BUILD_TIME="${BUILD_TIME}" \
             --build-arg BUILD_GO_VERSION="${BUILD_GO_VERSION}" \
             --build-arg BUILD_PATH="${DOCKER_BUILD_PATH}" \
             --build-arg VERSION_PATH="${VERSION_PATH}" \
              --build-arg MAIN_FILE="${MAIN_FILE}" \
             -t "${IMG}" .

docker-push: ## Push docker image with the manager.
	docker push ${IMG}

# PLATFORMS defines the target platforms for  the manager image be build to provide support to multiple
# architectures. (i.e. make docker-buildx IMG=myregistry/mypoperator:0.0.1). To use this option you need to:
# - able to use docker buildx . More info: https://docs.docker.com/build/buildx/
# - have enable BuildKit, More info: https://docs.docker.com/develop/develop-images/build_enhancements/
# - be able to push the image for your registry (i.e. if you do not inform a valid value via IMG=<myregistry/image:<tag>> then the export will fail)
# To properly provided solutions that supports more than one platform you should use this option.
PLATFORMS ?= linux/arm64,linux/amd64,linux/s390x,linux/ppc64le
.PHONY: docker-buildx
docker-buildx: test ## Build and push docker image for the manager for cross-platform support
	# copy existing Dockerfile and insert --platform=${BUILDPLATFORM} into Dockerfile.cross, and preserve the original Dockerfile
	#sed -e '1 s/\(^FROM\)/FROM --platform=\$$\{BUILDPLATFORM\}/; t' -e ' 1,// s//FROM --platform=\$$\{BUILDPLATFORM\}/' Dockerfile > Dockerfile.cross
	- docker buildx create --name project-v3-builder
	docker buildx use project-v3-builder
	- docker buildx build --load --platform $(PLATFORMS) --build-arg BUILD_BRANCH="${BUILD_BRANCH}" \
             --build-arg BUILD_COMMIT="${BUILD_COMMIT}" \
             --build-arg BUILD_TIME="${BUILD_TIME}" \
             --build-arg BUILD_GO_VERSION="${BUILD_GO_VERSION}" \
             --build-arg BUILD_PATH="${DOCKER_BUILD_PATH}" \
             --build-arg VERSION_PATH="${VERSION_PATH}" \
              --build-arg MAIN_FILE="${MAIN_FILE}" \
             -t "${IMG}" -f Dockerfile .
	- docker buildx rm project-v3-builder
	#rm Dockerfile.cross