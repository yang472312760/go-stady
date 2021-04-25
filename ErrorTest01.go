package main

import (
	"errors"
	"fmt"
)

func Divide(a, b float64) (result float64, err error) {
	if b == 0 {
		result = 0.0
		err = errors.New("runtime error: divide by zero")
		return
	}

	result = a / b
	err = nil
	return
}

func main() {
	r, err := Divide(10.0, 0)
	if err != nil {
		fmt.Println(err) //错误处理 runtime error: divide by zero
	} else {
		fmt.Println(r) // 使用返回值
	}
}
