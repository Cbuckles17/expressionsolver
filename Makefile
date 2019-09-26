BINARY_NAME=expressionsolver
WINDOWS=$(BINARY_NAME)_windows.exe
LINUX=$(BINARY_NAME)_linux
DARWIN=$(BINARY_NAME)_darwin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get -u
RELEASE_DIR=release

.PHONY: all clean

all: build

build: windows linux darwin

run-windows: windows
	$(RELEASE_DIR)//$(WINDOWS)

run-linux: linux
	$(RELEASE_DIR)//$(LINUX)

run-darwin: darwin
	$(RELEASE_DIR)/$(DARWIN)

windows: $(WINDOWS)

linux: $(LINUX)

darwin: $(DARWIN)

$(WINDOWS):
	mkdir -p $(RELEASE_DIR)
	env GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(RELEASE_DIR)/$(WINDOWS)

$(LINUX):
	mkdir -p $(RELEASE_DIR)
	env GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(RELEASE_DIR)/$(LINUX)

$(DARWIN):
	mkdir -p $(RELEASE_DIR)
	env GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(RELEASE_DIR)/$(DARWIN)

clean:
	rm -f $(RELEASE_DIR)/$(WINDOWS) $(RELEASE_DIR)/$(LINUX) $(RELEASE_DIR)/$(DARWIN)