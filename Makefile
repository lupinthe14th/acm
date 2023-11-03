# vi: ft=make

VERSION = $(shell git describe --always --tags --abbrev=0)
CURRENT = $(VERSION)
NEXT = $(shell svu next)
MESSAGE = $(shell git log $(CURRENT).. --pretty=format:"%s")
DEFAULT_BRANCH = $(shell git symbolic-ref refs/remotes/origin/HEAD | sed 's@^refs/remotes/origin/@@')

docker-buildx:
	VERSION=$(VERSION) docker buildx bake -f compose.yml --push

docker-buildx-print:
	VERSION=$(VERSION) docker buildx bake -f compose.yml --print

bumper:
	git switch $(DEFAULT_BRANCH)
	git fetch -p -P
	git pull
	git tag $(NEXT) -m "$(MESSAGE)"
	git push --tags
