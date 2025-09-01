#!/bin/bash

#基本
echo "Hello World"

#变量输出
name="Jay"
echo "HELLO $name!"

#不带引号的输出：可实现，但为避免意外最好带
echo 不带引号的输出内容

#不换行：git-bash环境下似乎有时不生效
echo -n "不换行"
echo "不换行内容"

#启动转移
echo "无转义\n效果"
echo -e "有转义\n效果"

#常用转义序列
echo -e "\n常用转义序列-----:"
#\\\\n在bash双引号处理下变成:\\n
#\\n在echo -e转义处理下变成:\n  因而打印出来
echo -e "\\\\n换行 \n换行内容示例\n"

echo -e "\\\\t水平制表符"
echo -e "\t水平制表符示例\n"

echo -e "\\\\v垂直制表符"
echo -e "\v垂直制表符示例\n"

echo -e "\\\\b退格"
echo -e " \b退格示例\n" #输出里吃掉了一个空格

echo -e "\\\\r回车"
#光标回到本行最前，不换行，因而后续输出会覆盖
echo -e "回车示例原内容\r回车示例覆盖内容"

#高级用法
echo -e "\n高级用法-----:"

#输出到文件
#单>应该是覆盖式的写法
echo "This will be saved to file" > output.txt
cat output.txt
#追加内容
#双>>应该是追加式的写法
echo "Additional line" >> output.txt
cat output.txt

#彩色输出
#前景色：30(黑)、31(红)、32(绿)、33(黄)、34(蓝)、35(紫)、36(青)、37(白)
#背景se：40-47 对应上述颜色
#\033为ESC字符（ASCII码27的八进制表示）标准开头
#m为属性的终止符号
#[为ANSI转义序列开始符号，ANSI序列是控制终端显示行为的标准化代码系统
echo -e "\033[31mRed Text\033[0m"
#开始转义，背景42，前景31，内容，结束（利用重置样式当做结束）
echo -e "\033[42;31mGreen Background with Red Text\033[0m"
#如果要这样分别实现，需要在每次更换前重置样式
echo -e "\033[33;45m新增\033[0m  \033[36;41m内容\033[0m  \033[31;44m详情\033[0m"

#使用命令替换输出结果
echo "今天是 $(date)"

#菜单设计
echo -e "\n\033[1mSystem Menu\033[0m"
echo "1. Check disk space"
echo "2. List running processes"
echo "3. Show system info"
echo -n "Please enter your choice [1-3]: "

#进度条设计
echo -n "Progress: ["
for i in {1..5}; do
    echo -n "#"
    sleep 0.1
done
echo "] Done!"

#在变量引用方面，双引号是推荐的形式，不添加可能因为变量含空格出问题
var="Hello World"
echo $var    # 可能有问题，如果变量包含空格
echo "$var"  