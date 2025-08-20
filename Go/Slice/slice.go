// 切片
// 切片是对数组的抽象，go数组长度不可改，因而通过切片作为一种灵活、强悍的内置类型动态数组。
// 切片长度不固定，且可以追加元素

package main

import (
	"fmt"
)

/* 定义格式
var slice_name []type
*/

/*	通过make()函数定义切片
var slice_name []type=make([]type,length,capacity)

slice_name:=make([]type,length,capacity)
*/

func main() {

	// 直接初始化
	s := []int{1, 2, 3} //初始cap=len=3
	fmt.Println(s)

	// 引用数组初始化
	var arr = [7]int{1, 2, 3, 4, 5, 6, 7}
	s1 := arr[:]   //arr[start_index,end_index]
	s2 := arr[1:4] //index 1 到 index 3,类似于py的逻辑
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println()

	// 通过切片初始化
	s3 := s1[3:5]
	fmt.Println(s3)

	// make
	s4 := make([]int, 2, 3)
	fmt.Println(s4)

	// len()长度函数
	// 切片是可索引，因而可用此函数获取长度
	fmt.Println("s1长度", len(s1))
	//cap()容量函数
	fmt.Println("s1容量", cap(s1))
	fmt.Println()

	// 空切片
	var numbers []int
	if numbers == nil {
		fmt.Println("numbers为空切片")
		fmt.Printf("len=%v cap=%v slice=%v\n", len(numbers), cap(numbers), numbers)
		fmt.Println()
	}

	// append()增加内容
	var slice_sum []int
	printSlice(slice_sum)
	// 允许对空切片进行追加
	slice_sum = append(slice_sum, 0)
	printSlice(slice_sum)
	// 追加一个元素
	slice_sum = append(slice_sum, 1)
	printSlice(slice_sum)
	// 追加多个元素
	slice_sum = append(slice_sum, 2, 3, 4)
	printSlice(slice_sum)

	// 基于切片容量创建新的切片
	slice_sum2 := make([]int, len(slice_sum), (cap(slice_sum) * 2))
	// 复制小的到大的
	copy(slice_sum2, slice_sum)
	printSlice(slice_sum2)
}

// 打印切片函数
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
