package main

import "fmt"

func TestA01() {
	fmt.Println("func TestA01()")
}

func TestB01(x int) {
	var a [10]int
	a[x] = 222 //x值为11时，数组越界
}

func TestC01() {
	fmt.Println("func TestC()")
}

func main() {
	TestA01()
	TestB01(11)
	TestC01()
}
