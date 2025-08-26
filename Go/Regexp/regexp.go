// 正则表达式
package main

import (
	"fmt"
	"regexp"
)

//regexp库常见函数与方法
/*
	Compile:编译，返回*Regexp对象和err
	MustCompile:编译，失败时直接panic
	MatchString:检查字符串是否匹配正则表达式
	FindString:查找匹配的字符串，返回第一个
	FindAllString:查找匹配字符串，返回所有
	ReplaceAllString:替换匹配的字符串
	Split:根据正则表达式分割字符串
*/

/*	基本语法
	.					匹配任意单个字符
	*					匹配前面的字符0或多次
	+					匹配前面的字符1或多次
	?					匹配前面的字符0或1次
	\d=[0-9]			匹配数字
	\w=[a-zA-Z0-9_]		匹配字母、数字、下划线
	\s					匹配空白字符（空格、制表符、换行符）
	[]					匹配括号内任意字符
	^					开头
	s					结尾
*/

func main() {

	// 匹配模式
	pattern := `^[a-zA-Z0-9]+$`
	// 编译正则表达式，生成正则表达式对象
	regex := regexp.MustCompile(pattern)
	str := "Hello123"
	// 判断匹配
	if regex.MatchString(str) {
		fmt.Println("字符串匹配正则表达式")
	} else {
		fmt.Println("字符串不匹配正则表达式")
	}

	// 查找所有匹配字符
	patternFA := `\d+`
	regexFA := regexp.MustCompile(patternFA)

	strFA := "我有3个苹果和5个香蕉"
	// 结果:-1表全部返回、0表返回nil、n>0表需要返回的结果数量
	matches := regexFA.FindAllString(strFA, -1)
	fmt.Println("找到的数字：", matches)

	// 替换匹配的字符串
	patternRP := `\s+`
	regexRP := regexp.MustCompile(patternRP)
	strRP := "Hello      World       !"
	// 把多空格替换成单空格
	result := regexRP.ReplaceAllString(strRP, " ")
	fmt.Println("替换后字符串：", result)

	// 分割字符串
	patternSP := `,`
	regexSP := regexp.MustCompile(patternSP)
	strSP := "apple,banna,orange"
	parts := regexSP.Split(strSP, -1)
	fmt.Println("分割后的字符串：", parts)
}
