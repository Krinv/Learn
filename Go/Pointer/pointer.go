// 指针
package main

import "fmt"

func main() {
	var a int = 20
	var p *int

	p = &a

	fmt.Println("a的地址", &a)
	fmt.Println("p存储内容", p)
	fmt.Print("p访问内容", *p)

	// 空指针
	var ptr *int
	fmt.Println("空指针为", ptr)

	// 指针数组
	c := []int{10, 100, 200}
	var ptrs [3]*int

	for i := 0; i < 3; i++ {
		ptrs[i] = &c[i]
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("c[%d]=%d\n", i, *ptrs[i])
	}

	//指向指针的变量
	var d int = 0
	var p2 *int = &d
	var p22 **int = &p2

	fmt.Println("d=", d)
	fmt.Println("*p2=", *p2)
	fmt.Println("**p22=", **p22)
}
