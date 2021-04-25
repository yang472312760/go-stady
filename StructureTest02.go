package main

import "fmt"

type Student01 struct {
	id int

	name string

	sex byte

	age int

	addr string
}

func main() {

	var s5 *Student01 = &Student01{id: 5, name: "yang", sex: 'm', age: 15, addr: "JL"}

	s6 := &Student01{4, "rocco", 'm', 3, "sh"}

	s7 := &s5

	fmt.Println(s6)
	fmt.Println(s7)

}
