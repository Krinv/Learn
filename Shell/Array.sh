#!/bin/bash

#数组
#定义数组
echo -e "\n定义数组------:"
array1=(1 2 3 4)
array1=(
    1
    2
    3
    4
)
array1[1]=1
echo "${array1[@]}"

#获取长度
echo -e "\n获取数组长度------:"
length=${#array1[@]}
echo $length
length=${#array1[*]}
echo $length

#获取单个元素长度
echo -e "\n获取单个元素长度------:"
length=${#array1[2]}
echo $length

#由于没有多行注释符，因而可以使用多种方式作为注释使用

#多行#号
#

# 通过空命令结合Here Document语法实现
# ：为空命令符，不执行任何操作，返回退出状态码0
# <<是here文档的起始符，将后续内容作为输入到前面的指令
:<<'COMMENT'    #单引号是为了防止here文档中的变量或命令被展开
多行注释内容
COMMENT

:<<!
注释内容
注释内容
!

# 关联数组：即使用字符串、整数作为下标来访问的数组
echo -e "\n声明关联数组------:"

#declare -A array_name  -A为关联数组的选项
#统一设置
declare -A site=(
    ["google"]="google.com"
    ["baidu"]="baidu.com"
    ["taobao"]="taobao.com"
)
echo "${site[@]}"

# 先声明再设置键值
declare -A site2
site2["google"]="google.com"
site2["baidu"]="baidu.com"
site2["taobao"]="taobao.com"
echo "${site2[@]}"