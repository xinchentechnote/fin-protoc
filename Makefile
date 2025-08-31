BIN_DIR := bin
LIB_DIR := lib

TARGET := fin-protoc
SHARED_LIB := libpacketdsl

UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

all: build

ifeq ($(OS),Windows_NT)
    # Windows
    ANTLR_JAR := $(subst \,/,${USERPROFILE})/software/antlr-4.13.2-complete.jar
else
    # Unix/Linux/Mac
    ANTLR_JAR := ${HOME}/software/antlr-4.13.2-complete.jar
endif

VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT := $(shell git rev-parse HEAD)
LDFLAGS := -X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME) -X main.gitCommit=$(GIT_COMMIT)

gen: 
	java -jar "$(ANTLR_JAR)" -Dlanguage=Go -no-listener -visitor -package gen -o internal grammar/PacketDsl.g4

build: gen dirs main-build shared-build

main-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/$(TARGET) ./cmd/

print-system:
	@echo "Detected system: UNAME_S = $(UNAME_S)"

shared-build: print-system
ifeq ($(OS),Windows_NT) 
	@echo "Building for native Windows..."
	go build -ldflags "$(LDFLAGS)" -buildmode=c-shared -o lib/libpacketdsl.dll ./cmd/
else ifneq (,$(findstring MSYS_NT,$(UNAME_S)))
	@echo "Building for Windows (MSYS2)..."
	go build -ldflags "$(LDFLAGS)" -buildmode=c-shared -o lib/libpacketdsl.dll ./cmd/
else ifeq ($(UNAME_S),Linux)
	@echo "Building for Linux..."
	go build -ldflags "$(LDFLAGS)" -buildmode=c-shared -o lib/libpacketdsl.so ./cmd/
else ifeq ($(UNAME_S),Darwin)
	@echo "Building for macOS..."
	go build -ldflags "$(LDFLAGS)" -buildmode=c-shared -o lib/libpacketdsl.dylib ./cmd/
else
	@echo "Unsupported system: $(UNAME_S)"
	exit 1
endif

package:shared-build main-build
	mkdir -p release

ifeq ($(OS),Windows_NT) 
	# Windows
	mkdir -p temp-pack/fin-protoc-$(VERSION)-windows-amd64/bin
	mkdir -p temp-pack/fin-protoc-$(VERSION)-windows-amd64/lib
	mkdir -p temp-pack/fin-protoc-$(VERSION)-windows-amd64/include
	cp $(BIN_DIR)/fin-protoc.exe temp-pack/fin-protoc-$(VERSION)-windows-amd64/bin/fin-protoc.exe
	cp $(LIB_DIR)/libpacketdsl.dll temp-pack/fin-protoc-$(VERSION)-windows-amd64/lib/
	cp $(LIB_DIR)/libpacketdsl.h temp-pack/fin-protoc-$(VERSION)-windows-amd64/include/
	cd temp-pack && zip -r ../release/fin-protoc-$(VERSION)-windows-amd64.zip \
		fin-protoc-$(VERSION)-windows-amd64/bin/fin-protoc.exe \
		fin-protoc-$(VERSION)-windows-amd64/lib/libpacketdsl.dll \
		fin-protoc-$(VERSION)-windows-amd64/include/libpacketdsl.h
else ifneq (,$(findstring MSYS_NT,$(UNAME_S)))
	@echo "Building for Windows (MSYS2)..."
	mkdir -p temp-pack/fin-protoc-$(VERSION)-windows-amd64/bin
	mkdir -p temp-pack/fin-protoc-$(VERSION)-windows-amd64/lib
	mkdir -p temp-pack/fin-protoc-$(VERSION)-windows-amd64/include
	go build -ldflags "$# Windows
	cp $(BIN_DIR)/fin-protoc.exe temp-pack/fin-protoc-$(VERSION)-windows-amd64/bin/fin-protoc.exe
	cp $(LIB_DIR)/libpacketdsl.dll temp-pack/fin-protoc-$(VERSION)-windows-amd64/lib/
	cp $(LIB_DIR)/libpacketdsl.h temp-pack/fin-protoc-$(VERSION)-windows-amd64/include/
	cd temp-pack && zip -r ../release/fin-protoc-$(VERSION)-windows-amd64.zip \
		fin-protoc-$(VERSION)-windows-amd64/bin/fin-protoc.exe \
		fin-protoc-$(VERSION)-windows-amd64/lib/libpacketdsl.dll \
		fin-protoc-$(VERSION)-windows-amd64/include/libpacketdsl.h
else ifeq ($(UNAME_S),Linux)
	# Linux
	mkdir -p temp-pack/fin-protoc-$(VERSION)-linux-amd64/bin
	mkdir -p temp-pack/fin-protoc-$(VERSION)-linux-amd64/lib
	mkdir -p temp-pack/fin-protoc-$(VERSION)-linux-amd64/include
	cp $(BIN_DIR)/fin-protoc temp-pack/fin-protoc-$(VERSION)-linux-amd64/bin/fin-protoc
	cp $(LIB_DIR)/libpacketdsl.so temp-pack/fin-protoc-$(VERSION)-linux-amd64/lib/
	cp $(LIB_DIR)/libpacketdsl.h temp-pack/fin-protoc-$(VERSION)-linux-amd64/include/
	tar -czvf release/fin-protoc-$(VERSION)-linux-amd64.tar.gz \
		-C temp-pack fin-protoc-$(VERSION)-linux-amd64/bin/fin-protoc \
		fin-protoc-$(VERSION)-linux-amd64/lib/libpacketdsl.so \
		fin-protoc-$(VERSION)-linux-amd64/include/libpacketdsl.h
else ifeq ($(UNAME_S),Darwin)
	# macOS
	mkdir -p temp-pack/fin-protoc-$(VERSION)-darwin-amd64/bin
	mkdir -p temp-pack/fin-protoc-$(VERSION)-darwin-amd64/lib
	mkdir -p temp-pack/fin-protoc-$(VERSION)-darwin-amd64/include
	cp $(BIN_DIR)/fin-protoc temp-pack/fin-protoc-$(VERSION)-darwin-amd64/bin/fin-protoc
	cp $(LIB_DIR)/libpacketdsl.dylib temp-pack/fin-protoc-$(VERSION)-darwin-amd64/lib/
	cp $(LIB_DIR)/libpacketdsl.h temp-pack/fin-protoc-$(VERSION)-darwin-amd64/include/
	tar -czvf release/fin-protoc-$(VERSION)-darwin-amd64.tar.gz \
		-C temp-pack fin-protoc-$(VERSION)-darwin-amd64/bin/fin-protoc \
		fin-protoc-$(VERSION)-darwin-amd64/lib/libpacketdsl.dylib \
		fin-protoc-$(VERSION)-darwin-amd64/include/libpacketdsl.h
else
	@echo "Unsupported system: $(UNAME_S)"
	exit 1
endif	

	rm -rf temp-pack

dirs:
	mkdir -p $(BIN_DIR) $(LIB_DIR)

run: build
	$(BIN_DIR)/$(TARGET) format -f internal/parser/testdata/sample_binary.dsl

test:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	rm -rf $(BIN_DIR) $(LIB_DIR) coverage.*

setup:
	sudo apt-get update && sudo apt-get install -y gcc-mingw-w64-x86-64

.PHONY: all build main-build shared-build cross-build dirs run test clean setup