#!/bin/bash

#if
echo -e "if相关内容补充------:"
#写成一行的if
if [ $(ps -ef | grep -c "ssh") -gt 1 ];then echo "true";else echo "false";fi

#多选项情况
#注意：在[]中大于-gt、小于-lt   在(())大于>、小于<
if [ 1 = 2 ]
then 
    echo "1=2"
elif [ 2 = 2 ]
then
    echo "2=2"
fi


#for
echo -e "\nfor相关内容补充------:"
for loop in 1 2 3 4 5
do
    echo "The value is : $loop"
done

echo -e "\nwhile相关内容补充------:"
int=1
while (($int<=5))
do
    echo $int
    let "int++"
done

echo -e "\nwhile用于输入内容------:"
echo -n '输入你最喜欢的网站名'
#read 定义变量FILM

while read FILM
do
    echo "是的！$FILM 是个好网站"
    break
done

#util
echo -e "\nutil相关内容------:"
a=0
until [ ! $a -lt 10 ]
do
    echo $a
    let "a++"
done

#case
echo -e "\ncase相关内容------:"
echo "输入一个数字"
read num

case $num in
    1) echo "你选择1"
    ;;
    2|3|4) echo "你选择2、3、4"
    ;;
    *) echo "你并未输入1、2、3、4"
    ;;
esac
