// 递归
// 递归函数的基本组成：基准条件（终止条件）+递归条件（将问题分解为子问题）

package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

// 阶乘函数
func factorial(n int) int {
	// 基准
	if n == 0 {
		return 1
	}

	// 递归
	return n * factorial(n-1)
}
