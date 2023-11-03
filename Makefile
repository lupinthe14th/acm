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
	git switch $(DEFAULT_BRANCH) || { echo "Failed to switch branch"; exit 1; }
	git fetch -p -P || { echo "Failed to fetch from remote repository"; exit 1; }
	git pull || { echo "Failed to pull the latest code"; exit 1; }
	git tag $(NEXT) -m "$(MESSAGE)" || { echo "Failed to create a new tag"; exit 1; }
	git push --tags || { echo "Failed to push tags"; exit 1; }
