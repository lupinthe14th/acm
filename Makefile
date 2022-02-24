# vi: ft=make

VERSION ?= $(shell git describe --always --tags --abbrev=0)

docker-buildx:
	VERSION=$(VERSION) docker buildx bake -f compose.yml --push

docker-buildx-print:
	VERSION=$(VERSION) docker buildx bake -f compose.yml --print
