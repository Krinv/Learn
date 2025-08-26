// 类型断言
package main

import "fmt"

// 类型断言：用于检查接口值的实际类型的机制
/*	基本语法
	value,ok:=interfaceValue.(Type)

	value：断言成功后的具体类型的值
*/

func main() {
	var i interface{} = "Hello,Go!"

	// 尝试将i断言为string
	s, ok := i.(string)
	if ok {
		fmt.Println("断言字符串成功：", s)
	} else {
		fmt.Println("断言字符串失败")
	}

	// 尝试将i断言为int
	n, ok := i.(int)
	if ok {
		fmt.Println("断言整型成功：", n)
	} else {
		fmt.Println("断言整型失败")
	}

	/*	直接引发panic的断言形式
		value:=interfaceValue.(Type)

		n:=i.(int)
	*/
}

// 多类型选择
func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("整数", v)
	case string:
		fmt.Println("字符串", v)
	default:
		fmt.Println("未知类型")
	}
}
