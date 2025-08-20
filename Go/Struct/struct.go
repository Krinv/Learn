// 结构体
package main

import "fmt"

/*	定义格式
	type struct_variable_type struct{
		member definition,
		member definition,
		...
	}
*/

type Books struct {
	title   string
	author  string
	subject string
	book_id string
}

func main() {
	// 创建结构体:	1
	fmt.Println(Books{"Go", "Lok", "Te", "664488"})

	// 实例化
	var book1 Books

	// 描述
	book1.title = "GO语言"
	book1.author = "L"
	book1.subject = "教程"
	book1.book_id = "666"

	// 访问成员
	fmt.Println("boo1 title:", book1.title)
	fmt.Println("boo1 author:", book1.author)
	fmt.Println("boo1 subject:", book1.subject)
	fmt.Println("boo1 book_id:", book1.book_id)
	fmt.Println()

	// 使用结构体参数
	printBook(book1)

	// 结构体指针
	var struct_pointer *Books
	struct_pointer = &book1
	fmt.Println("结构体指针:", struct_pointer.title)
}

// 结构体参数的函数
func printBook(book Books) {
	fmt.Println("使用结构体参数")
	fmt.Println("boo1 title:", book.title)
	fmt.Println("boo1 author:", book.author)
	fmt.Println("boo1 subject:", book.subject)
	fmt.Println("boo1 book_id:", book.book_id)
	fmt.Println()
}
