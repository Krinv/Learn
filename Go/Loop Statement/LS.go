// 循环语句
package main

import "fmt"

func main() {
	// 无限循环语句

	/*
		for true{
			fmt.Println("无限循环")
		}
	*/

	// for循环：形式1
	sum1 := 0
	// for init;condition;post{}
	for i := 0; i <= 3; i++ {
		sum1 += i
	}

	// for循环：形式2
	sum2 := 0
	for sum2 < 5 {
		sum2++
	}

	// for循环：形式3
	sum3 := 0
	for { //无限循环直到符合条件
		sum3++
		if sum3 > 3 {
			break
		}
	}

	fmt.Printf("sum1=%v sum2=%v sum3=%v", sum1, sum2, sum3)

	// 迭代循环：针对字符串、数组、切片等
	strings := []string{"aaa", "bbb"}
	//索引、键值
	for i, s := range strings {
		fmt.Println(i, s)
	}

	// map例
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	// 读取key、value
	for key, value := range map1 {
		fmt.Printf("key is %d\nvalue is %.2f\n", key, value)
	}
	// 仅读key
	for key := range map1 {
		fmt.Printf("仅key is %d\n", key)
	}
	// 仅读value：利用抛弃值
	for _, value := range map1 {
		fmt.Printf("仅value is %.2f\n", value)
	}

	// 跳转语句,利用标签实现
	goto LOOP
LOOP:
	fmt.Println("jump")

}
