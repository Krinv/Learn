#!/bin/bash

:<<'COMMENT'



COMMENT

#函数定义
echo -e "\n函数------:"
demoFun(){
    echo "第一个函数"
}

demoFun

#带return的函数
funWithReturn(){
    echo "输入第一个值"
    read aNum
    echo "输入第二个值"
    read bNum
    echo "两个值分别为${aNum},${bNum}"
    return $(($aNum+$bNum))
}

#注意return只能返回0-255的整数，因此最好在函数中echo输出

funWithReturn
#$?取上个命令退出时得到的值
echo "两数之和为$? !"

#传递参数的函数
funWithParam(){
    echo "第一个参数为$1!"
    echo "第二个参数为$2!"
    echo "第十个参数为${10}!"
    echo "总共有$#个参数"
    echo "字符串输出所有参数:$*"
}

#shell中直接传递
echo -e "\nShell指令：funWithParam 1 2 3 4 5 6 7 8 9 34 73"
funWithParam 1 2 3 4 5 6 7 8 9 34 73