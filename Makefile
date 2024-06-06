#Variables

BINARY_NAME = coreutils
VERSION = $(shell git describe --tags --abbrev=0)
BUILD_DIR = build

#Targets
all: clean build rename

clean:
	rm -rf $(BUILD_DIR)

build:
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-darwin-amd64
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-windows-amd64.exe

rename:
	mv $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-linux-amd64 $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64
	mv $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-darwin-amd64 $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64
	mv $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-windows-amd64.exe $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe

.PHONY: all clean build rename
