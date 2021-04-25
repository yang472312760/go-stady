package main

import "fmt"

type MyInt int

type s string

func (a MyInt) Add(b MyInt) MyInt {
	return a + b
}

func Add01(a, b MyInt) MyInt {
	return a + b
}

func main() {

	var a MyInt = 1

	var b MyInt = 2

	fmt.Println("a.Add(b) = ", a.Add(b))

	fmt.Println("Add(a, b) = ", Add01(a, b))

}
