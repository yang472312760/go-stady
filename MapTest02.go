package main

import "fmt"

func main() {

	m1 := map[int]string{1: "yang", 2: "cheng", 3: "lin"}
	//迭代遍历1，第一个返回值是key，第二个返回值是value
	for k, v := range m1 {

		fmt.Printf("%d ----> %s\n", k, v)

	}

	//删除key值为3的map
	delete(m1, 3)

	for k, v := range m1 {
		fmt.Printf("%d ----> %s\n", k, v)
	}

}
