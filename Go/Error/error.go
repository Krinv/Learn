// 错误处理
// 采用显式返回错误的方式，非传统异常处理机制
package main

// 标准库定义error接口
/*
	// 实现Error()方法都可以作为错误
	type error interface{
		// 返回一个描述错误的字符串
		Error() string
	}
*/

// fmt包提供了对于错误的格式化输出支持
/*
	%v	默认格式
	%+v	如果支持，显示详细的错误信息
	%s	作为字符串输出
*/

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("This is an error")
	fmt.Println(err)

	// 验证错误信息函数
	value1, err1 := Squrt(1)
	fmt.Printf("value:%f,error:%s\n", value1, err1)
	value2, err2 := Squrt(-1)
	fmt.Printf("value:%f,error:%s\n", value2, err2)

	// 错误接口应用例
	value3, err3 := divide(10, 0)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(value3)
	}

	err4 := findItem(1)
	// error.Is:检查某个错误是否是特定的错误或由该错误包装
	// 检查err4是否由ErrNotFound包装
	if errors.Is(err4, ErrNotFound) {
		fmt.Println("Item not found")
	} else {
		fmt.Println("Other error:", err4)
	}

	// errors.as:将错误转换为特定类型以便进一步处理
	// 检查err5是否可以转换为MyError，若可以则将错误值err5赋值给myErr
	err5 := getError()
	var myErr *MyError
	if errors.As(err5, &myErr) {
		fmt.Printf("Custom error	Code:%d,Msg:%s\n", myErr.Code, myErr.Msg)
	}

	// panic、recover
	fmt.Println("Starting program...")
	safeFunction()
	fmt.Println("Program continued after panic")
}

// 定义返回错误信息的函数
func Squrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math:square root of negative number")
	}

	return f, errors.New("no errors")
}

// 错误接口应用
// 除法错误结构
type DivideError struct {
	Dividend int
	Divisor  int
}

// 应用错误接口:通常使用指针接收者
func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}

	return a / b, nil
}

// error.Is:检查某个错误是否是特定的错误或由该错误包装
var ErrNotFound = errors.New("not found")

// 返回error接口类型。因为fmt.Errorf是返回这个
func findItem(id int) error {
	// %w 使用动词包装（wrap）了ErrNotFound错误
	// wrap：允许一个错误包含另一个错误
	// 即创建了这个新错误，包含了ErrNotFoun的这个错误
	return fmt.Errorf("database error: %w", ErrNotFound)
}

// errors.as:将错误转换为特定类型以便进一步处理
type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	// sprintf是格式化字符串并返回结果，不直接打印在控制台
	return fmt.Sprintf("Code:%d,Msg:%s\n", e.Code, e.Msg)
}

func getError() error {
	// 实例化了一个MyError
	return &MyError{Code: 404, Msg: "Not Found"}
}

// panic：处理不可恢复的错误
// recover：用于从panic中恢复
func safeFunction() {
	// defer用于延迟执行一个函数，直到包含它的函数执行完成
	// 即使发生panic，defer语句也会被执行
	// 多个defer按后进先出LIFO顺序执行
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	// 触发panic，停止正常的执行流程
	panic("something went wrong")

	/* 运行流程
	调用safeFunction，设置defer恢复点
	触发panic
	异常处理，执行已注册的defer，recover（）捕获panic的值
	恢复，打印信息
	*/
}
