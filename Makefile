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
	go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/$(TARGET) ./cmd/

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

cross-build: gen dirs
	# Linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 \
	go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/fin-protoc-linux-amd64 ./cmd/
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 \
	go build -ldflags "$(LDFLAGS)" -buildmode=c-shared -o $(LIB_DIR)/libpacketdsl-linux-amd64.so ./cmd/

	# Windows
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
	CC="x86_64-w64-mingw32-gcc" \
	go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/fin-protoc-windows-amd64.exe ./cmd/
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
	CC="x86_64-w64-mingw32-gcc" \
	go build -ldflags "$(LDFLAGS)" -buildmode=c-shared -o $(LIB_DIR)/libpacketdsl-windows-amd64.dll ./cmd/

	# macOS
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 \
	go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/fin-protoc-darwin-amd64 ./cmd/
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 \
	go build -ldflags "$(LDFLAGS)" -buildmode=c-shared -o $(LIB_DIR)/libpacketdsl-darwin-amd64.dylib ./cmd/

package: cross-build
	mkdir -p release temp-pack/{bin,lib,include}
	
	# Linux
	cp $(BIN_DIR)/fin-protoc-linux-amd64 temp-pack/bin/fin-protoc
	cp $(LIB_DIR)/libpacketdsl-linux-amd64.so temp-pack/lib/
	cp $(LIB_DIR)/packetdsl.h temp-pack/include/
	tar -czvf release/fin-protoc-$(VERSION)-linux-amd64.tar.gz \
		-C temp-pack bin/fin-protoc \
		lib/libpacketdsl-linux-amd64.so \
		include/packetdsl.h
	
	# Windows
	cp $(BIN_DIR)/fin-protoc-windows-amd64.exe temp-pack/bin/fin-protoc.exe
	cp $(LIB_DIR)/libpacketdsl-windows-amd64.dll temp-pack/lib/
	cp $(LIB_DIR)/packetdsl.h temp-pack/include/
	cd temp-pack && zip -r ../release/fin-protoc-$(VERSION)-windows-amd64.zip \
		bin/fin-protoc.exe \
		lib/libpacketdsl-windows-amd64.dll \
		include/packetdsl.h
	
	# macOS
	cp $(BIN_DIR)/fin-protoc-darwin-amd64 temp-pack/bin/fin-protoc
	cp $(LIB_DIR)/libpacketdsl-darwin-amd64.dylib temp-pack/lib/
	cp $(LIB_DIR)/packetdsl.h temp-pack/include/
	tar -czvf release/fin-protoc-$(VERSION)-darwin-amd64.tar.gz \
		-C temp-pack bin/fin-protoc \
		lib/libpacketdsl-darwin-amd64.dylib \
		include/packetdsl.h
	
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