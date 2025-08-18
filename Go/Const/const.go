// 常量
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 显示定义方法
	const a string = "abc"

	// 隐式定义方法
	const b = "abc"

	// 多重赋值
	const c, d, e = 1, 12.5, false

	fmt.Println(a, b, c, d, e)

	// 常量用作枚举
	const (
		unknown = 0
		female  = 1
		male    = 2
	)

	// 内置函数实现常量表达式计算
	const (
		f = "abc"
		g = len(a)
		i = unsafe.Sizeof(a)
	)
	fmt.Println(f, g, i)

	// 特殊常量iota：可被编辑器修改的常量
	// 在const出现时，将被重置为0（const内部第一行前）
	// const内部每增加一行常量声明，则iota计数一次，因而可认为iota为const语块中的行索引
	const (
		j    = iota
		k    = iota
		l, m = iota, iota //可见是行索引，而非变量索引，每当iota在新的一行被使用时，会自增1
	)
	// 简写形式
	const (
		o = iota
		p
		q
	)
	fmt.Println(j, k, l, m)
	fmt.Println(o, p, q)

	// iota综合例
	const (
		r = iota //0
		s        //1
		t        //2
		u = "ha" //独立值"ha" iota=3
		v        //跟随上面"ha" iota=4
		w = 100  //独立值100,iota=5
		x        //跟随上面100,iota=6
		y = iota //7，恢复计数
		z        //8
	)
	fmt.Println(r, s, t, u, v, w, x, y, z)

	// 移位例子
	const (
		a1 = 1 << iota //a1=1且左移0位，得1
		a2 = 3 << iota //a2=3=b11且左移1位，110=6
		a3             //a3=3=b11且左移2位，1100=12
		a4             //a4=3=b11且左移3位，11000=24
	)
	fmt.Println(a1, a2, a3, a4)
}
