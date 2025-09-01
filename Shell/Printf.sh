#!/bin/bash

# 带变量的输出
echo -e "常规转义输出------:"
name="Alice"
printf "Hello, %s\n" "$name"

# 整数
printf "十进制Decimal: %d\n十六进制Hex: %x\n八进制Octal: %o\n" 255 255 255

# 浮点数
printf "Float: %f\nScientific: %e\n" 3.14159 3.14159

# 字符串
printf "Name: %s\n" "Bob"

# 字符
printf "First letter: %c\n" "A"

#格式化控制
echo -e "格式化控制------:"
printf "%10s\n%-10s\n" "right" "left"
printf "%04d\n" "23"

#多参数:成倍入参写入即可实现
printf "%-10s %5d %8.2f\n" "Apple" 5 2.5 "Orange" 3 1.75

#颜色输出
RED="\033[0;31m"
GREEN="\033[0;32m"
NOCOLOR="\033[0m"

printf "${RED}Error:${NOCOLOR}Something went wrong\n"
printf "${GREEN}Success:${NOCOLOR} Operation completed\n"
