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
ifeq ($(UNAME_S),Linux)
	@echo "Building Linux shared library..."
	go build -buildmode=c-shared -o $(LIB_DIR)/$(SHARED_LIB).so ./cmd/
endif
ifeq ($(UNAME_S),Windows_NT)
	@echo "Building Windows shared library..."
	go build -buildmode=c-shared -o $(LIB_DIR)/$(SHARED_LIB).dll ./cmd/
endif

# 创建必要的目录
dirs:
	mkdir -p $(BIN_DIR) $(LIB_DIR)

# 运行程序
run: build
	$(BIN_DIR)/$(TARGET) format -f internal/parser/testdata/need_format.dsl

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

.PHONY: all build main-build shared-build dirs run test clean setup