# 构建目录
BIN_DIR := bin
LIB_DIR := lib

# 目标名称
TARGET := fin-protoc
SHARED_LIB := libpacketdsl

# 检测操作系统
UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

# 默认目标
all: build

# 构建主程序和共享库
build: dirs main-build shared-build
# 仅构建主程序
main-build:
	go build -o $(BIN_DIR)/$(TARGET) ./cmd/

print-system:
	@echo "Detected system: UNAME_S = $(UNAME_S)"

# 共享库构建规则
shared-build: print-system
ifeq ($(OS),Windows_NT) 
	@echo "Building for native Windows..."
	go build -buildmode=c-shared -o lib/libpacketdsl.dll ./cmd/
else ifneq (,$(findstring MSYS_NT,$(UNAME_S)))
	@echo "Building for Windows (MSYS2)..."
	go build -buildmode=c-shared -o lib/libpacketdsl.dll ./cmd/
else ifeq ($(UNAME_S),Linux)
	@echo "Building for Linux..."
	go build -buildmode=c-shared -o lib/libpacketdsl.so ./cmd/
# else ifeq ($(UNAME_S),Darwin)
# 	@echo "Building for macOS..."
# 	go build -buildmode=c-shared -o lib/libpacketdsl.dylib ./cmd/
else
	@echo "Unsupported system: $(UNAME_S)"
	exit 1
endif

shared-build-osxcross: print-system
	# Linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 \
	go build -buildmode=c-shared -o $(LIB_DIR)/$(SHARED_LIB).so ./cmd/

	# Windows (需要安装MinGW)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
	CC="x86_64-w64-mingw32-gcc" \
	go build -buildmode=c-shared -o $(LIB_DIR)/$(SHARED_LIB).dll ./cmd/


# 创建必要的目录
dirs:
	mkdir -p $(BIN_DIR) $(LIB_DIR)

# 运行程序
run: build
	$(BIN_DIR)/$(TARGET) format -f internal/parser/testdata/sample_binary.dsl

# 运行测试
test:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# 清理
clean:
	rm -rf $(BIN_DIR) $(LIB_DIR) coverage.*
	
# 安装依赖工具
setup:
	# 安装MinGW (Windows交叉编译)
	sudo apt-get update && sudo apt-get install -y gcc-mingw-w64-x86-64

.PHONY: all build main-build shared-build dirs run test clean setup shared-build-osxcross