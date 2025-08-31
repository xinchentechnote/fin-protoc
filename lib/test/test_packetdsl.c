#include <stdio.h>
#include <stdlib.h>
#include "libpacketdsl.h"  // 包含生成的头文件

int main() {
    // 测试1: 正常输入
    char* input1 = 
    "options {\n"
    "    string_pre_fix_len_type   =  u16;\n"
    "     repeat_pre_fix_size_type  = u16;\n"
    "}\n"
    "\n"
    "packet SampleBinary {\n"
    "      uint16 MsgType  `消息类型`,\n"
    "    u16  BodyLenght `消息体长度`,\n"
    "    match MsgType {\n"
    "        1 :  Logon,\n"
    "        2 :  Logout,\n"
    "        3  : Heartbeat,\n"
    "        4  : RiskControlRequest,\n"
    "        5 :   RiskControlResponse,\n"
    "    }\n\n"
    "}";
    
    printf("=== Test 1: Valid DSL Input ===\n");
    char* result1 = FormatPacketDslExport(input1);
    printf("Input:\n%s\n", input1);
    printf("Output:\n%s\n", result1);
    free(result1);
    
    // 测试2: 空输入
    printf("\n=== Test 2: Empty Input ===\n");
    char* input2 = "";
    char* result2 = FormatPacketDslExport(input2);
    printf("Input: \"%s\"\n", input2);
    printf("Output: \"%s\"\n", result2);
    free(result2);
    
    // 测试3: 错误输入
    printf("\n=== Test 3: Invalid DSL Input ===\n");
    char* input3 = "invalid dsl string { missing bracket";
    char* result3 = FormatPacketDslExport(input3);
    printf("Input: %s\n", input3);
    printf("Output: %s\n", result3);
    free(result3);
    
    return 0;
}