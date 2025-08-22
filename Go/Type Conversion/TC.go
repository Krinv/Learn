// 类型转化
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 17
	var b int = 5
	var mean float32

	mean = float32(a) / float32(b)

	fmt.Printf("mean:%f\n", mean)

	// 字符串转整数
	var str string = "10"
	var num int
	num, _ = strconv.Atoi(str) //返回：转换后的整型值、可能发生的错误
	fmt.Printf("str为%s,转换整数后为%d\n", str, num)

	// 转化失败例
	var str1 string = "ABC"
	num1, error := strconv.Atoi(str1)
	fmt.Printf("str为%s,转换整数后为%d,信息%s\n", str1, num1, error)

	//整数转化为字符串
	num2 := 123
	str2 := strconv.Itoa(num2)
	fmt.Printf("num2为%d,转换为字符串后为%s\n", num2, str2)

	// 字符串解析浮点数
	str3 := "3.14"
	num3, error3 := strconv.ParseFloat(str3, 64)
	fmt.Printf("str3为%s,转换为浮点数为%f,信息%v\n", str3, num3, error3)

	// 浮点数转字符串
	fmt.Println()
	num4 := 3.444
	//要转换的数：num4；使用格式：f，浮点数格式；保留位数：2；原始浮点数位数：64
	str4 := strconv.FormatFloat(num4, 'f', 2, 64)
	fmt.Printf("浮点数为%f,转化为字符串为%v\n", num4, str4)

	// 接口类型转化：类型断言
	/*
		value.(type) value为接口类型的变量，type为要转化的类型
	*/
	fmt.Println()
	var i interface{} = "hello world" //定义接口变量赋值为字符串
	str, ok := i.(string)             //使用类型断言进行转换
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}

	// 接口类型转换：类型转换
	var w Writer = &StringWriter{} //定义w为Writer接口，创建StringWriter实例赋值w
	sw := w.(*StringWriter)        //通过类型断言将接口类型转换
	sw.str = "hello world"
	fmt.Println(sw.str)
	fmt.Println()

	// 空接口处理多类型值
	printValue(42)
	printValue("hello")
	printValue(3.14)
}

// 定义接口
type Writer interface {
	// 接收[]
	Write([]byte) (int, error)
}

// 实现Writer接口的结构体
type StringWriter struct {
	str string
}

// Write方法
func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

// 空接口类型
// 空接口可以持有任何类型的值，应用中常处理多种类型的值

func printValue(v interface{}) {
	// v.(type) 类型开关的专用语法，仅在switch中使用
	// 作用：获取接口变量v的动态类型（存储值的类型），而不是接口本身的静态类型（interface{}）
	switch v := v.(type) { //实际上是创建一个类型确定的新变量
	// 判断v为int
	case int:
		// 将v转换为int，保留原始值
		fmt.Println("Integer:", v) //因此调用时仍为v存储的值
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown")
	}
}
