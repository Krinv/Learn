// 运算符
package main

import "fmt"

func main() {
	// 算术运算符
	var a1 int = 21
	var a2 int = 10
	var a3 int

	a3 = a1 + a2
	fmt.Println("a1 + a2=", a3)
	a3 = a1 - a2
	fmt.Println("a1 - a2=", a3)
	a3 = a1 * a2
	fmt.Println("a1 * a2=", a3)
	a3 = a1 / a2
	fmt.Println("a1 / a2=", a3)
	a3 = a1 % a2
	fmt.Println("a1 % a2=", a3)
	// 自增减
	a1++
	fmt.Println("自增a1得", a1)
	a2--
	fmt.Println("自减a2得", a2)

	// 关系运算符
	var b1 int = 21
	var b2 int = 10

	if b1 == b2 {
		fmt.Println("b1等于b2")
	} else {
		fmt.Println("b1不等于b2")
	}

	if b1 != b2 {
		fmt.Println("不等于判定成功")
	}

	if b1 < b2 {
		fmt.Println("b1小于b2")
	} else {
		fmt.Println("b1不小于b2")
	}

	if b1 > b2 {
		fmt.Println("b1大于b2")
	} else {
		fmt.Println("b1不大于b2")
	}

	if b1 <= b2 {
		fmt.Println("b1小于等于b2")
	} else {
		fmt.Println("b1不小于等于b2")
	}

	if b1 >= b2 {
		fmt.Println("b1大于等于b2")
	} else {
		fmt.Println("b1不大于等于b2")
	}

	//逻辑运算符
	var c1 bool = true
	var c2 bool = false
	// AND
	if c1 && c2 {
		fmt.Println("c1&&c2为真")
	} else {
		fmt.Println("c1&&c2为假")
	}
	// OR
	if c1 || c2 {
		fmt.Println("c1||c2为真")
	} else {
		fmt.Println("c1||c2为假")
	}
	// NOT
	if !c2 {
		fmt.Println("!b为真")
	}

	// 位运算符
	// 60=0011 1100	13=0000 1101
	var d1, d2, d3 uint = 60, 13, 0
	// 按位与
	d3 = d1 & d2 //12=0000 1100
	fmt.Println("d1 & d2=", d3)
	// 按位或
	d3 = d1 | d2 //61=0011 1101
	fmt.Println("d1 | d2", d3)
	// 按位异或
	d3 = d1 ^ d2 //49=0011 0001
	fmt.Println("d1 ^ d2=", d3)
	// 左移
	d3 = d1 << 2 //240=1111 0000
	fmt.Println("d1<<2=", d3)
	// 算术右移
	d3 = d2 >> 2 //15=0000 1111
	fmt.Println("d2>>2=", d3)

	// 赋值运算符
	var e int = 21
	var f int
	// 赋值
	f = e
	fmt.Println("f=e得", f)

	f += e
	fmt.Println("f+=e得", f)

	f -= e
	fmt.Println("f-=e得", f)

	f *= e
	fmt.Println("f*=e得", f)

	f /= e
	fmt.Println("f/=e得", f)

	f = 200
	f <<= 2
	fmt.Println("f<<=2得", f)

	f >>= 2
	fmt.Println("f>>=2得", f)

	f &= 2
	fmt.Println("f&=2得", f)

	// 异或
	f ^= 2
	fmt.Println("f^=2得", f)

	f |= 2
	fmt.Println("f|=2得", f)

	// 其他运算符
	var ptr *int
	ptr = &f //取地址运算符

	fmt.Println("*ptr=", *ptr) //指针运算符
	fmt.Println("f=", f)
}
