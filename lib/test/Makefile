# 定义编译器
CC := gcc

# 定义目标
TARGET := test_packetdsl
LIBRARY := ../libpacketdsl.so

# 默认目标
all: build run

# 构建规则
build:
	$(CC) -o $(TARGET) test_packetdsl.c -L$(dir $(LIBRARY)) -I$(dir $(LIBRARY)) -lpacketdsl

# 运行规则
run: build
	@LD_LIBRARY_PATH=$$LD_LIBRARY_PATH:$(dir $(LIBRARY)) ./$(TARGET)

# 清理规则
clean:
	rm -f $(TARGET)

# 声明伪目标
.PHONY: all build run clean