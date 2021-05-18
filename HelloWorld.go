//1)、go语言以包作为单位管理
//2)、每个文件必须先声明包
//3)、程序必须有一个main包（重要）
package main

import "fmt"

//入口函数
func main() { //左括号必须和函数名同行
	//打印
	fmt.Println("HelloWorld") //go语言语句结尾没有分号

	var a string = "123456"

	fmt.Println(a)

	var b, c int = 10, 19

	d := 10
	

	fmt.Println(d)

	fmt.Println(b + c)
}
