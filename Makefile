VERSION := 0.1.39
NAME := dachsurl
DIST := $(NAME)-$(VERSION)
USER_NAME := practicet
REPO_NAME := $(USER_NAME)/$(NAME)
PACKAGE_LIST := $(shell go list ./...)

# $(NAME): coverage.out cmd/$(NAME)/main.go *.go
# $(NAME): coverage.out cmd/$(NAME)/main.go cmd/$(NAME)/generate_completion.go *.go
$(NAME): coverage.out
	go build -o $(NAME) cmd/$(NAME)/main.go cmd/$(NAME)/generate_completion.go
# 	go build -o $(NAME) cmd/$(NAME)/main.go

coverage.out:
	go test -covermode=count \
		-coverprofile=coverage.out $(PACKAGE_LIST)

docker: $(NAME)
#	docker build -t ghcr.io/$(REPO_NAME):$(VERSION) -t ghcr.io/$(REPO_NAME):latest .
	docker buildx build -t ghcr.io/$(REPO_NAME):$(VERSION) \
		-t ghcr.io/$(REPO_NAME):latest --platform=linux/arm64/v8,linux/amd64 --push .

# refer from https://pod.hatenablog.com/entry/2017/06/13/150342
define _createDist
	mkdir -p dist/$(1)_$(2)/$(DIST)
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME)$(3) cmd/$(NAME)/main.go cmd/$(NAME)/generate_completion.go
	cp -r README.md LICENSE dist/$(1)_$(2)/$(DIST)
#	cp -r docs/public dist/$(1)_$(2)/$(DIST)/docs
	tar cfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
endef

dist: $(NAME)
	@$(call _createDist,darwin,amd64,)
	@$(call _createDist,darwin,arm64,)
	@$(call _createDist,windows,amd64,.exe)
	@$(call _createDist,windows,arm64,.exe)
	@$(call _createDist,linux,amd64,)
	@$(call _createDist,linux,arm64,)

distclean: clean
	rm -rf dist

clean:
	rm -f $(NAME) coverage.out
	rm -rf completions cmd/dachsurl/completions
