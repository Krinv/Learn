#!/bin/bash

#内容重置
echo "_mbsetupuser console  Oct 31 17:35" > users.txt
echo "_mbsetupuser console  Oct 31 17:35" > users.txt
echo "_mbsetupuser console  Oct 31 17:35" > users.txt
echo "_mbsetupuser console  Oct 31 17:35" > users.txt
cat users.txt

#覆盖重定向
echo -e "\n输出覆盖重定向------:"
echo "覆盖内容" > users.txt
cat users.txt

#追加重定向
echo -e "\n输出追加重定向------:"
echo "追加内容" >> users.txt
cat users.txt

#输入重定向
echo -e "\n输入重定向"
#统计文件行数
wc -l users.txt
wc -l < users.txt   #此例不输出文件名，因为其仅从标准输入读取内容

#同时替换输入输出
cat < input.txt >>users.txt

:<<'COMMENT'
    在每个linux命令运行时，都会打开：
    1. 标准输入文件stdin：文件描述符为0，默认从stdin读取数据
    2. 标准输出文件stdout：文件描述符为1，默认从stdout输出数据
    3. 标准错误文件stderr：文件描述为2，unix程序会向stderr流中书写错误信息

    默认情况，command > file 将stdout重定向到file。command < file将stdin重定向到file

    #重定向stderr到文件指令格式
    command 2> file
    command 2>> file    追加到文件结尾
    
    #将stdout、stderr合并后重定向到file
    command >> file 2>&1

    #对stdin、stdout都重定向
    command < file1 > file2
COMMENT

#如果执行命令而不希望在屏幕显示结果
ls > /dev/null  #通过这个文件的内容都会丢弃

#希望屏蔽stdout、stdin
ls > /dev/null 2>&1

