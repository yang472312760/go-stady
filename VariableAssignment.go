package main

import "fmt"

func main() {

	var str string                                    // 声明一个字符串变量
	str = "abc"                                       // 字符串赋值
	ch := str[0]                                      // 取字符串的第一个字符
	fmt.Printf("str = %s, len = %d\n", str, len(str)) //内置的函数len()来取字符串的长度
	fmt.Printf("str[0] = %c, ch = %c\n", str[0], ch)

	//`(反引号)括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
	str2 := `hello
    mike \n \r测试
    `
	fmt.Println("str2 = ", str2)
	/*
	   str2 =  hello
	         mike \n \r测试
	*/

}
