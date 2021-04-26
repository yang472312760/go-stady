package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("子协程结束")

		fmt.Println("子协程正在运行……")

		c <- 666 //666发送到c
	}()

	num := <-c //从c中接收数据，并赋值给num

	fmt.Println("num = ", num)

	fmt.Println("main协程结束")

}
