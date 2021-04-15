package main

import "fmt"

func test(x int) {
	fmt.Println(100 / x) //x为0时，产生异常
}

func main() {
	defer fmt.Println("aaaaaaaa")
	defer fmt.Println("bbbbbbbb")

	defer test(0)

	defer fmt.Println("cccccccc")
	/*
	   运行结果：
	   cccccccc
	   bbbbbbbb
	   aaaaaaaa
	   panic: runtime error: integer divide by zero
	*/
}
