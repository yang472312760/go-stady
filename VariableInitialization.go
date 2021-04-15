package main

import (
	"fmt"
	"reflect"
)

func main() {

	var v1 int = 10 // 方式1
	var v2 = 10     // 方式2，编译器自动推导出v2的类型
	v3 := 10        // 方式3，编译器自动推导出v3的类型
	fmt.Println(v1 + v2)
	fmt.Println("v3 type is ", reflect.TypeOf(v3)) //v3 type is  int

	//出现在 := 左侧的变量不应该是已经被声明过，:=定义时必须初始化
	//var v4 int
	//v4 := 2 //err

}
