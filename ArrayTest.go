package main

import "fmt"

func main() {

	var a [10]int

	for i := 0; i < 10; i++ {
		a[i] = i + 1
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
	//range具有两个返回值，第一个返回值是元素的数组下标，第二个返回值是元素的值
	for i1, v1 := range a {
		fmt.Println("a[", i1, "]=", v1)
	}

}
