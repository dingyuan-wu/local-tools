BINARY_NAME ?= aws-account-switcher

.PHONY: build install clean

build:
	go build -o bin/$(BINARY_NAME) $(CMD_PATH)

install:
	go install $(CMD_PATH)

clean:
	rm -rf bin

