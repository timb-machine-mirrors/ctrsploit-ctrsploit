.PHONY: all shell local build

GITCOMMIT := $(shell git rev-parse --short HEAD || echo unsupported)
export GITCOMMIT

DOCKER_CONTAINER_NAME := $(if $(CONTAINER_NAME),--name $(CONTAINER_NAME),)
CTRSPLOIT_IMAGE := ctrsploit-dev
DOCKER_FLAGS := docker run --rm -ti $(DOCKER_CONTAINER_NAME) $(DOCKER_ENVS) $(DOCKER_MOUNT)

DOCKER_RUN_DOCKER := $(DOCKER_FLAGS) "$(CTRSPLOIT_IMAGE)"
DOCKERFILE := Dockerfile
BUILD_OPTS := ${BUILD_APT_MIRROR} ${DOCKER_BUILD_ARGS} ${DOCKER_BUILD_OPTS} -f "$(DOCKERFILE)"

LDFLAGS := ${LDFALGS} \
	-X github.com/ctrsploit/ctrsploit/version.Version=${VERSION} \
	-X github.com/ctrsploit/ctrsploit/version.GitCommit=${GITCOMMIT} \
	-X github.com/ctrsploit/ctrsploit/version.BuildTime=${BUILDTIME}

binary: bundle
	docker buildx bake binary --progress plain

bundle:
	mkdir -p bin/release

build-ctrsploit:
	./release.sh

build-image:
	docker buildx build $(BUILD_OPTS) --load -t "$(CTRSPLOIT_IMAGE)" --progress plain .

shell: build-image
	docker run --rm -ti -v $(CURDIR):/root/ctrsploit $(CTRSPLOIT_IMAGE) bash
