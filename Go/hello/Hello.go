// Hello World
package main //指明该文件属于哪个包,每个go程序必须要有一个main

import "fmt" //导入fmt包，内含格式化IO函数

//注意{不能放在单独的一行，否则会出错
func main() {
	fmt.Println("Hello, World!")

	//字符串连接
	fmt.Println("Google" + " Golang")
}

//Go中一行代表一个语句的结束。若多个语句在一行，需要人为用;分割

/*多行注释*/
