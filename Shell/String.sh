#!/bin/bash

#字符串
#单引号：
#   任何字符都会原样输出，单引号字符串中的变量是无效的
#   单引号字符串中不能出现单独的单引号，即使使用转义也不行，必须成对坐字符串拼接使用
str='a string'

#双引号:可以有变量、转义字符
your_name="John"
str="Hello,\"$your_name\"!\n"
echo -e $str

#拼接字符串
echo "拼接字符串"
your_name="John"
#双引号拼接
greeting="Hello,"$your_name" !"
greeting_1="hello,${your_name}!"
echo $greeting $greeting_1
#单引号拼接
greeting='Hello,'$your_name' !'
greeting_1='hello,${your_name}!'
echo $greeting $greeting_1

#获取长度
echo -e "\n获取字符串长度:"
str="abcd"
echo $str
#${#string}等价于${#string[0]}
echo "长度为：" ${#str}

#类似切片的提取子字符串
echo -e "\n提取子字符串:"
str="12345678"
#索引机制是类似的
echo ${str:1:4}
echo ${str::4}
echo ${str:1:}  #放空默认为索引0
echo ${str::}
echo ${str::-1} #支持负索引
echo ${str:1:7}

#查找子字符串
echo -e "\n查找子字符串："
str="Hello John! What are you doing right now?"
#expr index指令：查询str中i或o的位置。这里首先o出现在最早的hello，所以先返回o位置5
echo `expr index "$str" io`