package main

import "fmt"

func TestA1() {
	fmt.Println("func TestA()")
}

func TestB1() (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("internal error: %v", x)
		}
	}()
	panic("func TestB(): panic")
}

func TestC1() {
	fmt.Println("func TestC()")
}

func main() {
	TestA1()
	err := TestB1()
	fmt.Println(err)
	TestC1()
}
