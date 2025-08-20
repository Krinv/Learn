// 范围函数
// range()用于for中迭代数组、切片、通道、集合的元素。
// 于数组、切片中返回索引和value；于集合中返回key_value
package main

import "fmt"

func main() {

	// 数组、切片
	var pow = []int{1, 2, 4, 8, 16}
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}

	// 字符串:迭代时返回索引、Unicode代码点rune
	fmt.Println()
	for i, c := range "Hello" {
		fmt.Printf("index: %d char: %c\n", i, c)
	}

	//集合: map[index] type
	fmt.Println()
	map1 := make(map[int]float32)
	// 添加值
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for key, value := range map1 {
		fmt.Printf("key: %d  value: %f\n", key, value)
	}

	// 通道
	fmt.Println()
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
