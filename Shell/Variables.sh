#!/bin/bash

#部分有效的变量名，避免使用$等特殊字符
NAME="Jacy"
LIBRARY_PATH="/bin"

# 等号两侧避免使用空格
VALUE_NAME="value"

#除了显式赋值，可用语句给变量赋值
#循环当前文件夹所有文件
for file in *;do
    echo "处理文件：$file"
done

#echo默认不转义,-e开启转义
echo -e "\n"
echo $'\n'
printf "%s\n%s\n\n" "第一行" "第二行"

#使用$字符的实现一样的效果
for file in *;do
    echo "文件：$file"
done

#使用定义过的变量：使用$符号
YOUR_NAME="Jay"
echo $YOUR_NAME
echo ${YOUR_NAME}   #加入{}是为了更好识别变量边界

#如下例不用{}容易出错
SKILL="java"
echo "I am good at ${SKILL}Script"

#只读变量定义
URL="www.baidu.com"
readonly URL
echo $URL
# 运行URL="111"时会报错

#删除变量,注意unset不能删除只读变量
unset SKILL
echo $SKILL

#字符串变量定义
string1="Hello!"
string2='HELLO!'
echo $string1 $string2

#声明整数变量
typeset -i my_integer1=42
echo $my_integer1
declare -i my_integer2=43
echo $my_integer2

#数组
echo $'\n'
my_array=(1 2 3 4 5)
#@符号用于传递给脚本或函数所有参数的位置，每个参数作独立字符串处理，空格会保留
for element in "${my_array[@]}";do
    echo "$element"
done
#*符号传递时，元素内空格会被破坏。file name拆成file、name
echo $'\n'
#不加双引号时是分词处理
for element in ${my_array[*]};do
    echo "$element"
done
#加双引号时会整合成一个
for element in "${my_array[*]}";do
    echo "$element"
done

echo "不加引号导致的问题"
#不加引号导致的问题
echo "元素的错误拆分"
my_array=("file 1.txt" "file2.txt")
for element in ${my_array[@]};do
    echo "$element"
done
echo "通配符的错误展开"
my_array=("*.sh" "data")
for element in ${my_array[@]};do
    echo "$element"
done
echo "空字符串元素丢失"
my_array=("file1" " " "file3")
for element in ${my_array[@]};do
    echo "$element"
done

#关联数组操作
declare -A associative_array
associative_array["name"]="John"
associative_array["age"]=30
echo ${associative_array["name"]}
echo ${associative_array["age"]}

#环境变量
echo $PATH

#特殊变量
echo ${0} #脚本名
echo ${1} #脚本的参数
echo ${#} #传递给脚本的数量
echo $? #上个命令的退出状态

