// 文件处理
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

/*
	"os"	//文件处理核心库,创建、读写、删除等底层文件操作
	"io"	//通用接口，与文件、网络等数据源交互
	"bufio" //通过缓冲减少IO次数，适合频繁读写
	"path"	//跨平台兼容(win与linux的路径分隔符差异)
*/

func main() {

	// 使用os包进行文件创建
	file, err := os.Create("createfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //保证在函数结束时关闭
	log.Println("文件创建成功")

	// 打开文件
	// os.Open以只读模式打开
	file_open, err_open := os.Open("readfile.txt")
	if err_open != nil {
		log.Fatal(err_open)
	}
	defer file_open.Close()
	log.Println("文件打开成功")

	// 逐行读取文件
	scanner := bufio.NewScanner(file_open)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("文件逐行读取失败")
	} else {
		log.Println("文件逐行读取成功")
	}

	// 一次性读取文件
	// 直接是提取出了内容，而非打开着文件
	content, err_open1 := os.ReadFile("readfile.txt")
	if err_open1 != nil {
		fmt.Println("文件整体读取失败")
	} else {
		fmt.Println(string(content))
		log.Println("文件整体读取成功")
	}

	// 常规写入
	// fileWrite, errWrite := os.OpenFile("writefile.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	fileWrite, errWrite := os.Create("writefile.txt")
	//os.O_WRONLY写入	os.O_CREATE创建文件 	os.O_APPEND追加内容和权限0644
	if errWrite != nil {
		log.Fatal(errWrite)
	} else {
		// 直接写入字符串
		fileWrite.WriteString("直接写入的字符串内容\n")

		// 写入字节切片
		data := []byte("写入字节切片\n")
		fileWrite.Write(data)

		// 使用fmt写入
		fmt.Fprintf(fileWrite, "格式化写入：%d\n", 123)

		log.Println("文件常规写入成功")
	}
	defer fileWrite.Close()

	// 逐行写入文件
	writer := bufio.NewWriter(fileWrite)
	fmt.Fprintln(writer, "Hello World")
	writer.Flush()
	log.Println("文件逐行写入成功")

	// 一次性写入文件
	content2 := []byte("Hello!")

	// err_write := os.WriteFile("writefile.txt", content2, 0644) 这个写法是覆盖写法

	// 追加写入方法
	fileWrite1, errWrite1 := os.OpenFile("writefile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errWrite1 != nil {
		fmt.Println("文件整体写入失败")
	} else {
		if _, err := fileWrite1.Write(content2); err != nil {
			log.Println("文件整体写入失败")
		} else {
			log.Println("文件整体写入成功")
		}

	}
	defer fileWrite1.Close()

	// 删除文件操作
	fileDelete, errCreate := os.Create("deletefile.txt")
	if errCreate != nil {
		log.Println("创建删除用文件失败")
	}
	fileDelete.Close() //这里要及时关闭，否则会因为一个文件打开而无法进行后续删除

	errDelete := os.Remove("deletefile.txt")
	if errDelete != nil {
		log.Println("删除文件失败")
		fmt.Println(errDelete)
	} else {
		log.Println("删除文件成功")
	}

	// 获取文件信息
	fileInfo, errInfo := os.Stat("writefile.txt")
	if errInfo != nil {
		log.Fatal(err) //Fatal函数是直接退出程序
	}

	fmt.Println("文件名", fileInfo.Name())
	fmt.Println("文件大小", fileInfo.Size(), "字节")
	fmt.Println("权限", fileInfo.Mode())
	fmt.Println("最后修改时间", fileInfo.ModTime())
	fmt.Println("是否为目录", fileInfo.IsDir())

	// 检查文件是否存在
	// errState包含后面函数无法正常运行的各种错误原因
	// os.IsNotExist()函数判断错误信息是否由文件不存在导致
	if _, errState := os.Stat("write"); os.IsNotExist(errState) {
		log.Println("文件不存在")
	} else {
		log.Println("文件存在")
	}

	// 重命名文件
	errRename := os.Rename("old.txt", "new.txt")
	if errRename != nil {
		log.Println(errRename)
	} else {
		log.Println("重命名成功new")
	}

	errRename = os.Rename("new.txt", "old.txt")
	if errRename != nil {
		log.Println(errRename)
	} else {
		log.Println("重命名成功old")
	}

	// 创建目录
	errMkdir := os.Mkdir("newdir", 0755)
	if errMkdir != nil {
		log.Println(errMkdir)
	} else {
		log.Println("单级目录创建成功")
	}
	// 递归创建目录
	errMkdir = os.MkdirAll("new/new/new", 0755)
	if errMkdir != nil {
		log.Println(errMkdir)
	} else {
		log.Println("递归创建目录成功")
	}

	// 读取目录内容
	entries, errEn := os.ReadDir(".") //当前目录
	if err != nil {
		log.Println(errEn)
	}
	// 抛弃值为索引，entry为目录内容
	for _, entry := range entries {
		info, _ := entry.Info()
		// 格式化：左对齐占20字符宽度、右对齐占8字符宽度、默认格式输出
		fmt.Printf("%-20s %8d %v\n",
			entry.Name(),
			info.Size(),
			info.ModTime().Format("2006-01-02 15:04:05"), //时间格式化
		)
	}

	// 删除空目录
	// err := os.Remove("emptydir")

	// // 递归删除目录及其内容
	// err = os.RemoveAll("path/to/dir")

	// 文件复制
	// 源文件
	srcFile, errSrc := os.Open("source.txt")
	if errSrc != nil {
		log.Println(errSrc)
	}
	defer srcFile.Close()

	// 目标文件
	dstFile, errDist := os.Create("destination.txt")
	if errDist != nil {
		log.Println(errDist)
	}
	defer dstFile.Close()

	// 复制操作
	bytesCopied, errByte := io.Copy(dstFile, srcFile)
	if errByte != nil {
		log.Println("文件复制失败")
	} else {
		log.Println("文件复制成功共复制", bytesCopied, "字节")
	}

	// 临时文件
	// "":创建文件的目录，放空为系统默认临时目录
	// 在 Unix-like 系统上通常是 /tmp
	// 在 Windows 上通常是 %TEMP% 或 %TMP% 环境变量指定的目录
	// "example-*.txt"：文件名模型，通配符*会被随机字符串替代以保证文件名唯一性
	tmpFile, errTmp := os.CreateTemp("", "example-*.txt")
	if errTmp != nil {
		log.Println("创建临时文件失败")
		log.Println(errTmp)
	}
	defer os.Remove(tmpFile.Name())
	fmt.Println("临时文件：", tmpFile.Name())

	// 创建临时目录
	tmpDir, errTmpDir := os.MkdirTemp("", "example-*")
	if errTmpDir != nil {
		log.Println("创建临时目录失败")
		log.Println(errTmpDir)
	}
	defer os.Remove(tmpDir)
	fmt.Println("临时目录：", tmpDir)
}
