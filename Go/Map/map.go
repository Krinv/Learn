// 集合
// 无序的键值对集合。通过key来快速检索数据。
// Map可迭代，但遍历Map时返回的键值对顺序是不确定的。键不存在时，返回该类型的零值。
// Map是引用类型，因而将一个Map传递给另一个函数或变量时，都指向同一个底层数据结构，对Map的修改会影响所有引用它的变量
package main

import "fmt"

func main() {
	// 使用make定义
	/*	格式
		make_variable:=make(map[keyType]valueType,initialCapacity)
	*/
	m := make(map[string]int, 10)
	fmt.Println("m为", m)

	// 使用字面量创建
	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
	}
	fmt.Println("m1为", m1)

	// 获取值
	v1 := m1["apple"]
	fmt.Println("v1为", v1)
	v2, v3 := m1["pear"] //从map中获取值时会返回两个：value和key是否存在的bool值
	fmt.Println("v2为", v2, "v3为", v3)

	// 修改元素
	m1["apple"] = 5
	fmt.Println("apple", m1["apple"])

	// 获取长度
	fmt.Println("len", len(m1))

	// 遍历Map
	for k, v := range m1 {
		fmt.Printf("key=%s value=%d\n", k, v)
	}

	// 删除键值对
	delete(m1, "banana")
	fmt.Println("删除后m1", m1)
}
