/*
	函数定义格式
	func function_name([parameter list])  [return types]{

	}
*/

// 函数
package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2
	var ret = max(a, b)

	var a1 string = "111"
	var b1 string = "222"
	a1, b1 = swap(a1, b1)

	fmt.Printf("最大值为%v\n交换后值为a1=%v b1=%v\n\n", ret, a1, b1)

	// 值传递验证
	var x int = 1
	var y int = 2
	fmt.Printf("值传递交换前的x=%v,y=%v\n", x, y)
	swap1(x, y)
	fmt.Printf("值传递交换后的x=%v,y=%v\n\n", x, y)

	// 引用传递实现
	fmt.Printf("引用传递交换前的x=%v,y=%v\n", x, y)
	swap2(&x, &y)
	fmt.Printf("引用传递交换后的x=%v,y=%v\n", x, y)

	// 可以使用函数作为实参
}

// 返回最大值函数
func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

// 多值返回函数定义
func swap(x, y string) (string, string) {
	return y, x
}

// 默认情况下采用值传递调用参数：即类c语言，变化的是复制的参数
func swap1(x, y int) {
	var temp int

	temp = x
	x = y
	y = temp
}

// 利用引用传递的方法实现改变入参值
func swap2(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
