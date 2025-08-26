// 条件操作
package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2

	// if
	if a > b {
		fmt.Println("a>b")
	} else {
		fmt.Println("a不大于b")
	}

	// switch
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		a = 222
	}

	// type switch	判断某个interface变量中实际存储的变量类型
	var x interface{}

	switch x.(type) {
	case nil:
		fmt.Printf("类型为:nil\n")
	case int:
		fmt.Printf("类型为:int\n")
	default:
		fmt.Println("未知类型")
	}

	// fallthrough会强制执行后面的case语句，并且不会判断下条case表达式是否为true
	// 空条件，相当于true
	switch {
	case false: //不会进入
		fmt.Println("1.false")
		fallthrough
	case true: //首先进入
		fmt.Println("2.true")
		fallthrough //接着执行下个打印
	case false:
		fmt.Println("3.false")
		fallthrough //接着执行下个打印
	case true:
		fmt.Println("4.true") //因为fallthrough仅执行一条，接下来的5不会再执行，于此结束
	case false:
		fmt.Println("5.false")
		fallthrough
	default:
		fmt.Println("6.默认case")
	}

}
