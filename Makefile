PACKAGE_LIST := $(shell go list ./...)
VERSION := 1.0.0
NAME := dachsurl
DIST := $(NAME)-$(VERSION)

# $(warning PACKAGE_LIST = $(PACKAGE_LIST))

all: build test

build:
	go build -o $(NAME) $(PACKAGE_LIST)

test:
	gofmt -l -s .
	go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)

distclean: clean
	rm -rf dist
	rm -f coverage.out

clean:
	rm -rf dachsurl
