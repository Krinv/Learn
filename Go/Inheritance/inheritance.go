// 继承
package main

/*
	Go并非传统面向对象编程语言，无类和继承概念，而是通过struct、interfce、composition实现类似功能
*/

import "fmt"

// composition：组合是Go中代码复用主要方式
// 父结构体
type Animal struct {
	Name string
}

// 父结构体方法
func (a *Animal) Speak() {
	fmt.Println(a.Name, "say hello!")
}

// 子结构体
type Dog struct {
	Animal //嵌入父结构
	Breed  string
}

// 接口
type Speaker interface {
	Speak() //方法
}

func main() {
	// 实例
	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	dog.Speak() //直接调用父结构方法
	fmt.Println("Breed:", dog.Breed)

	// 直接通过接口调用方法
	var speaker Speaker
	speaker = &dog //将实例赋给接口，通过接口调用
	speaker.Speak()

	// 涉及重写方法的例子
	fmt.Println()
	v := Vechicle{Brand: "Toyota"}
	c := Car{
		// 注意这里的初始化方式
		Vechicle: Vechicle{Brand: "Honda"},
		Model:    "Civic",
	}

	v.Start()
	c.Start()

	// 亦可调用父原方法
	c.Vechicle.Start()
}

// 另一个涉及重写方法的例子
// 基类
type Vechicle struct {
	Brand string
}

// 基类方法
func (v *Vechicle) Start() {
	fmt.Println(v.Brand, "started!")
}

// 派生类
type Car struct {
	Vechicle
	Model string
}

// 重写方法
func (c *Car) Start() {
	fmt.Println(c.Brand, c.Model, "car started!")
}
