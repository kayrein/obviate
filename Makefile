BINARY_NAME=obviate
PACKAGE_PATH=github.com/kayrein/obviate

VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -X main.version=$(VERSION)

.PHONY: build install clean

build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) .

install:
	go install -ldflags "$(LDFLAGS)" $(PACKAGE_PATH)

clean:
	rm -f $(BINARY_NAME)
