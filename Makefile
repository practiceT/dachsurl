PACKAGE_LIST := $(shell go list ./...)
VERSION := 1.0.0
NAME := dachsurl
DIST := $(NAME)-$(VERSION)

# $(warning PACKAGE_LIST = $(PACKAGE_LIST))

all:
	:

dachsurl: coverage.out
	go build -o $(NAME) $(PACKAGE_LIST)

coverage.out:
	go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)

clean:
	rm -rf dachsurl dist coverage.out
