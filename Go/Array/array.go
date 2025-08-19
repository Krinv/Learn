// 数组
package main

import "fmt"

/*	创建格式
	var array_name [size]data_type
*/

func main() {
	// 声明数组:自动根据类型初始化
	var balance [10]float32
	fmt.Println(balance[0])

	// 通过初始化列表来初始化数组元素
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 短声明初始化
	num := [2]int{1, 2}
	fmt.Println(num)

	// 未知长度数组
	var bal1 = [...]float32{1, 2, 3}
	fmt.Println(bal1)
	bal2 := [...]float32{1, 2}
	fmt.Println(bal2)

	// 指定下标初始化（对于限定长度）
	var num3 = [3]int{1: 2, 0: 1}
	fmt.Println(num3)

	// 多维数组
	// var array_name [x][y] variable_type
	value := [][]int{}
	row1 := []int{1, 1, 1}
	row2 := []int{2, 2, 2}

	value = append(value, row1)
	value = append(value, row2)

	fmt.Println(value)

	// 初始化多维数组，使用{}
	a := [3][3]int{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3}} //括号不能单独一行，如果要独行，最后一个{}加上,号
	fmt.Println(a)

	// 创建内部不同长度的多维数组
	animals := [][]string{}

	row3 := []string{"fish", "cat"}
	row4 := []string{"f", "c", "dog"}

	animals = append(animals, row3)
	animals = append(animals, row4)

	fmt.Println(animals)

	// 验证以数组为参数时，不改变原数组的内容，需要通过指针实现原数组的修改
	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Origin:", myArray)

	modifyArray(myArray)
	fmt.Println("Modifyarray:", myArray)

	modifyArrayWithPointet(&myArray)
	fmt.Println("Pointer:", myArray)

}

// 接受数组的函数
func modifyArray(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

// 接受数组指针的函数
func modifyArrayWithPointet(arr *[5]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
}
