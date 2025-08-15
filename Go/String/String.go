package main

import "fmt"

//格式化参数生成字符串
func main() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"

	//Sprintf根据格式化参数生成格式化字符串并返回
	//第一个参数为格式化字符串
	//后面为入参
	var target_url = fmt.Sprintf(url, stockcode, enddate)

	//Printf根据格式化参数生成格式化字符串并写入标准输出

	fmt.Println(target_url)
}
