package main

import "fmt"

func Test01() (int, string) { //方式1
	return 123, "abc"
}

func Test02() (a int, str string) { //方式2, 给返回值命名
	a = 456
	str = "def"
	return
}

func main() {
	v1, v2 := Test01() //函数调用
	_, v3 := Test02()  //函数调用， 第一个返回值丢弃
	v4, _ := Test02()  //函数调用， 第二个返回值丢弃
	fmt.Printf("v1 = %d, v2 = %s, v3 = %s, v4 = %d\n", v1, v2, v3, v4)
}
