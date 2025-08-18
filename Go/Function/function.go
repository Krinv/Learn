/*
	函数定义格式
	func function_name([parameter list])  [return types]{

	}
*/

// 函数
package main

import (
	"fmt"
	"math"
)

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
	fmt.Printf("引用传递交换后的x=%v,y=%v\n\n", x, y)

	// 可以使用函数作为另一个函数的实参
	// 声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	// 使用函数
	fmt.Println(getSquareRoot(4))
	fmt.Println()

	// 匿名函数验证
	next := getSequence()
	fmt.Println("序号", next())
	fmt.Println("序号", next())
	fmt.Println("序号", next())

	next2 := getSequence() //重新创建则重新计数
	fmt.Println("重新序号", next2())
	fmt.Println()

	// 方法验证
	var c Circle
	c.radius = 10
	fmt.Println("圆的面积：", c.getArea())
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

// 匿名函数(闭包)：无函数名，一般用在函数内部定义函数，或者作为函数参数
func getSequence() func() int { //返回值类型为func() int，即无参且返回int类型的一个函数
	i := 0
	return func() int {
		i++ //闭包特性：访问并修改外部作用域的变量
		return i
	}
}

// 方法：一个函数，内含有接受者信息。
// 接受者：命名类型或结构体类型的一个值或一个指针
// 所有给定类型的方法，属于该类型的方法集
/*
	方法格式：
	func (variable_name variable_data_type) function_name() return_type {
	}
*/
type Circle struct {
	radius float32
}

// 接受者为Circle结构体的一个方法
func (c Circle) getArea() float32 {
	return 3.14 * c.radius * c.radius
}
