package main

import "fmt"

//通过循环实现1+2+3……+100
func Test1() int {
	i := 1
	sum := 0
	for i = 1; i <= 100; i++ {
		sum += i
	}

	return sum
}

//通过递归实现1+2+3……+100
func Test2(num int) int {
	if num == 1 {
		return 1
	}

	return num + Test2(num-1) //函数调用本身
}

//通过递归实现1+2+3……+100
func Test3(num int) int {
	if num == 100 {
		return 100
	}

	return num + Test3(num+1) //函数调用本身
}

func main() {

	fmt.Println(Test1())    //5050
	fmt.Println(Test2(100)) //5050
	fmt.Println(Test3(1))   //5050
}
