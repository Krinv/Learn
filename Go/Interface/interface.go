// 接口
// 用于定义行为的集合，通过描述类型必须实现的方法，规定类型的行为契约
// Go提供的接口数据类型，将所以有共性的方法定义在一起，任何类型只要实现了这些方法，即为实现了这个接口

/*
接口特点：

隐式定义：

	无关键字显式声明某类型实现了某接口
	一个类型实现了接口要求的所有方法，该类型就自动被认为实现了该接口

接口类型变量：

	可存储实现该接口的任意值
	包含：动态类型（存储实际值类型）、动态值（存储实际值）

零值接口：

	接口零值为nil
	未初始化接口变量的值为nil，且不包含任何类型或值

空接口：

	可以表示任何类型

接口的常见用法：

	多态|同一行为在不同对象上的不同实现形式：不同类型实现同一接口
	解耦|降低代码互相依赖程度：通过接口定义依赖关系，降低模块间的耦合
	泛化：空接口表示任意类型
*/
package main

import (
	"fmt"
	"math"
)

/* 定义格式
// 接口
type interface_name interface{
	method_name1 [return_type]
	method_name2 [return_type]
	......
}

// 结构体
type struct_name struct{
	......
}

// 实现接口方法
func (struct_name_variable struct_name)method1_name()[return_type]{

}
......
*/

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle 实现 Perimeter 接口
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 空接口
func printValue(val interface{}) {
	fmt.Printf("Value:%v 	Type:%T\n", val, val)
}

// type switch
func printType(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	}
}

func main() {
	c := Circle{Radius: 5}

	// 接口变量可以存储实现了接口的类型
	var s Shape = c
	fmt.Println("Area", s.Area())
	fmt.Println("Perimeter", s.Perimeter())
	fmt.Println()

	printValue(42)
	printValue("Hello")

	// 接口断言
	var i interface{} = "hello"
	str := i.(string)
	fmt.Println(str)
	fmt.Println()

	// 类型选择
	printType(42)
	printType("HELLO")
	fmt.Println()

	// 接口组合
	var rw ReadWriter = File{}
	fmt.Println(rw.Read())
	rw.Write("Hello Go")
}

// 接口组合
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

// 组合两个接口
type ReadWriter interface {
	Reader
	Writer
}

type File struct{}

// 对File结构应用接口
func (f File) Read() string {
	return "Reading data"
}
func (f File) Write(data string) {
	fmt.Println("Writing data:", data)
}
