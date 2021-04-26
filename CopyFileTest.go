package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	args := os.Args //获取用户输入的所有参数

	//如果用户没有输入,或参数个数不够,则调用该函数提示用户
	if args == nil || len(args) != 3 {
		fmt.Println("useage : xxx srcFile dstFile")
		return
	}

	srcPath := args[1] //获取输入的第一个参数
	dstPath := args[2] //获取输入的第二个参数

	fmt.Printf("srcPath = %s, dstPath = %s\n", srcPath, dstPath)

	if srcPath == dstPath {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	openFile, err := os.Open(srcPath) //打开源文件

	if err != nil {
		fmt.Println(err)
		return
	}

	createFile, err := os.Create(dstPath) //创建目的文件

	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, 1024) //切片缓冲区

	for {
		//从源文件读取内容，n为读取文件内容的长度
		read, err := openFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}

		if read == 0 {
			break
		}
		//切片截取
		bytes := buf[:read]
		//把读取的内容写入到目的文件
		createFile.Write(bytes)
	}

	//关闭文件
	openFile.Close()
	createFile.Close()
}
