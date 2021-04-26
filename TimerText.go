package main

import (
	"fmt"
	"time"
)

func main() {
	//创建定时器，2秒后，定时器就会向自己的C字节发送一个time.Time类型的元素值
	timer1 := time.NewTimer(time.Second * 2)

	t1 := time.Now()

	fmt.Printf("t1: %v\n", t1)

	t2 := <-timer1.C

	fmt.Printf("t2: %v\n", t2)

	//如果只是想单纯的等待的话，可以使用 time.Sleep 来实现
	timer2 := time.NewTimer(time.Second * 2)
	<-timer2.C
	fmt.Println("2s后")

	time.Sleep(time.Second * 2)
	fmt.Println("再一次2s后")

	<-time.After(time.Second * 2)
	fmt.Println("再再一次2s后")

	timer3 := time.NewTimer(time.Second)
	go func() {
		<-timer3.C
		fmt.Println("Timer 3 expired")
	}()

	stop := timer3.Stop() //停止定时器
	if stop {
		fmt.Println("Timer 3 stopped")
	}

	fmt.Println("before")
	timer4 := time.NewTimer(time.Second * 5) //原来设置3s
	timer4.Reset(time.Second * 1)            //重新设置时间
	<-timer4.C
	fmt.Println("after")

}
