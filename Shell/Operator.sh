#!/bin/bash

#原生bash不支持简单的数学计算，需要通过其他的命令实现
#expr是一款表达式计算工具

echo -e "\n两数相加------:"
#注意空格，这个其实算是入参，因而不能当做表达式
#目前``被认为是遗留的语法，不推荐使用
val=$(expr 2 + 2)
echo "两数之和为：$val"

#实例
a=10
b=20

#expr是外部进程，每次使用都要重新创建新进程
#当前bash以内置更高效、安全的算术运算语法,这个就不是入参，就是表达式形式
echo -e "更高效安全的算术运算------:"
val=$((a+b))
echo "a+b=$val"
val=$((a-b))
echo "a-b=$val"
val=$((a*b))
echo "a*b=$val"
val=$((a/b))
echo "a/b=$val"
val=$((a%b))
echo "a%b=$val"

#简单条件
#注意[]与变量间必须存在的空格
if [ $a == $b ]
then
    echo "a等于b"
fi

if [ $a != $b ]
then
    echo "a不等于b"
fi

#关系运算符
echo -e "\n关系运算符------:"

#-eq是否等于
if [ $a -eq $b ]
then
    echo "$a -eq $b : a等于b"
else
    echo "$a -eq $b : a不等于b"
fi

#-ne是否不等于
if [ $a -ne $b ]
then
    echo "$a -ne $b : a不等于b"
else
    echo "$a -ne $b : a等于b"
fi

#-gt是否大于
if [ $a -gt $b ]
then
    echo "$a -gt $b : a大于b"
else
    echo "$a -gt $b : a不大于b"
fi

#-lt是否小于
if [ $a -lt $b ]
then
    echo "$a -lt $b : a小于b"
else
    echo "$a -lt $b : a不小于b"
fi

#-ge是否大于等于
if [ $a -ge $b ]
then
    echo "$a -ge $b : a大于等于b"
else
    echo "$a -ge $b : a不大于等于b"
fi

#-le是否小于等于
if [ $a -le $b ]
then
    echo "$a -le $b : a小于等于b"
else
    echo "$a -le $b : a不小于等于b"
fi

#布尔
a=10
b=20
echo -e "\n布尔运算符------:"
#非运算
if [ $a != $b ]
then
    echo "$a != $b : a不等于b"
else
    echo "$a != $b : a等于b"
fi
#或运算
if [ $a -lt 100 -o $b -gt 5 ]
then
    echo "a小于100或者b大于5"
fi
#且运算
if [ $a -lt 100 -a $b -gt 5 ]
then
    echo "a小于100且b大于5"
else
    echo "a不小于100或b不大于5"
fi

#在逻辑判断中目前更推荐[p] || [q]
if [ $a -lt 100 ] || [ $b -gt 5 ]
then
    echo "a小于100或者b大于5"
fi

#在逻辑判断中目前更推荐 [p] && [q]
if [ $a -lt 100 ] && [ $b -gt 5 ]
then
    echo "a小于100且b大于5"
else
    echo "a不小于100或b不大于5"
fi

#字符串运算符
echo -e "\n字符串运算符------:"
a="abc"
b="efg"

#因为有取值符号，所以单符号=可以用于判断
if [ $a = $b ]
then
    echo "$a = $b : a=b"
else
    echo "$a = $b : a!=b"
fi

if [ $a != $b ]
then
    echo "$a != $b : a=b"
else
    echo "$a != $b : a!=b"
fi

#-z检测字符串长度是否为0
if [ -z $a ]
then
    echo "-z $a : 字符串长度为0"
else
    echo "-z $a : 字符串长度不为0"
fi

#-n检测字符串长度是否不为0,必须使用双引号将字符串包括
if [ -n "$a" ]
then
    echo "-n $a : 字符串长度不为0"
else
    echo "-n $a : 字符串长度为0"
fi

#直接引用字符串：判断字符串是否为空
if [ $a ]
then
    echo "$a : 字符串不为空"
else
    echo "$a : 字符串为空"
fi

#文件测试运算符
echo -e "\n文件测试运算符-----:"
file="./Test.sh"

if [ -b $file ]
then
    echo "-b $file:是块设备文件"
else
    echo "-b $file:不是块设备文件"
fi

if [ -c $file ]
then
    echo "-c $file:是字符设备文件"
else
    echo "-c $file:不是字符设备文件"
fi

if [ -d $file ]
then
    echo "-d $file:是目录"
else
    echo "-d $file:不是目录"
fi

if [ -f $file ]
then
    echo "-f $file:是普通文件"
else
    echo "-f $file:不是普通文件"
fi

#一般应用目录上，目录有sgid权限时，任何用户在该目录创建文件的属组会继承该目录的属组
if [ -g $file ]
then
    echo "-g $file:是设置了SGID位"
else
    echo "-g $file:不是设置了SGID位"
fi

#Stickybit:粘滞位，设置后仅文件所有者或root能删除与移动文件
if [ -k $file ]
then
    echo "-k $file:是设置了粘着位"
else
    echo "-k $file:不是设置了粘着位"
fi

if [ -p $file ]
then
    echo "-p $file:是有名管道"
else
    echo "-p $file:不是有名管道"
fi

#让普通用户临时拥有该文件属主的执行权限。仅能用在二进制可执行文件。而且suid权限只能设置在属主位置上
if [ -u $file ]
then
    echo "-u $file:是设置了SUID位"
else
    echo "-u $file:不是设置了SUID位"
fi

if [ -r $file ]
then
    echo "-r $file:是可读"
else
    echo "-r $file:不是可读"
fi

if [ -w $file ]
then
    echo "-w $file:是可写"
else
    echo "-w $file:不是可写"
fi

if [ -x $file ]
then
    echo "-x $file:是可执行"
else
    echo "-x $file:不是可执行"
fi

if [ -s $file ]
then
    echo "-s $file:是为空"
else
    echo "-s $file:不是为空"
fi

if [ -e $file ]
then
    echo "-e $file:是存在"
else
    echo "-e $file:不是存在"
fi

if [ -S $file ]
then
    echo "-S $file:是socket"
else
    echo "-S $file:不是socket"
fi

if [ -L $file ]
then
    echo "-L $file:是存在且是一个符号链接"
else
    echo "-L $file:不是存在或不是一个符号链接"
fi

#自增自减
echo -e "\n自增和自减操作的实现-----:"

#使用let
num=5
let num++
echo $num
#使用$(( ))、(( ))语法
((num++))
echo $num