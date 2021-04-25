package main

import "fmt"

func DeleteMapByKey(m map[int]string, key int) {

	delete(m, key) //删除key值为3的map

	for k, v := range m {
		fmt.Printf("len(m)=%d, %d ----> %s\n", len(m), k, v)
	}

}

func main() {

	m := map[int]string{1: "yang", 2: "cheng", 3: "lin"}

	DeleteMapByKey(m, 3)

	for k, v := range m {
		fmt.Printf("len(m)=%d, %d ----> %s\n", len(m), k, v)
	}

}
