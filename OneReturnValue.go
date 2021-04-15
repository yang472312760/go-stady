package main

import "fmt"

func Test001() int { //方式1
	return 123
}

//官方建议：最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差
func Test002() (value int) { //方式2, 给返回值命名
	value = 123
	return value
}

func Test003() (value int) { //方式3, 给返回值命名
	value = 123
	return
}

func main() {
	v1 := Test001() //函数调用
	v2 := Test002() //函数调用
	v3 := Test003() //函数调用
	fmt.Printf("v1 = %d, v2 = %d, v3 = %d\n", v1, v2, v3)
}
