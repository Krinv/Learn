package main

import "fmt"

// 全局变量声明一般采用因式分解关键字的写法
// 全局变量允许声明而不使用
// 局部变量不允许声明而不使用
var (
	n int
	o string
)

func main() {
	// 定义单一变量
	// 关键字+变量名+类型
	var a string = "root"
	fmt.Println(a)

	// 定义多个变量，用，分割
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 未初始化的变量默认值为0
	var d bool //即false
	var e string
	fmt.Println(d, e)

	// 未指定type的情况下会给默认对应类型
	var f = 'A' //推断为rune即int32的别名
	fmt.Println(f)

	// 显式优先：同时指定类型和初始化值，以显示类型为准
	var g float32 = 42 //定义为float32
	fmt.Println(g)

	//复杂类的默认值默认为nil，即未映射
	// pointer
	var h *int
	fmt.Println(h)
	// slice切片
	var i []int
	fmt.Println(i)
	// map映射
	var j map[string]int //键为string，值为int的映射
	fmt.Println(j)
	// channel通道
	var k chan int //传递int类型的通道
	fmt.Println(k)
	// function
	var l func(string) int //入参string 返回int
	fmt.Println(l)
	// interface
	var m error //error是一个内置接口
	fmt.Println(m)

	// 短声明方法:仅能在函数体内使用
	intVal := 1
	fmt.Println(intVal)

	// 格式化输出
	fmt.Printf("%v %v %v %q\n", i, f, b, a)
	fmt.Println(n, o)

	// 对于值类型而言，变量直接指向存在内存中的值
	// 使用=赋值时，其实将内存中的值进行了拷贝
	var p int = 5
	var q = p
	fmt.Println(p, q, &p, &q) //获取内存地址,通常存储在栈中

	// 便捷的交换值方法
	var a1, a2 int
	a1 = 1
	a2 = 2
	a2, a1 = a1, a2
	fmt.Printf("a1=%v,a2=%v", a1, a2)

	//因为go局部变量都要求使用到的特性产生了抛弃值的需求
	// 空白符_用于抛弃值，因为有时候用不上一个函数的所有返回值
	_, a2 = a1, a2 //_为一个只写变量，你得不到他的值
}
