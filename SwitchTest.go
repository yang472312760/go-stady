package main

import "fmt"

type Element01 interface {
}

type Person01 struct {
	name string
	age  int
}

func main() {

	list := make([]Element01, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person01{"yang", 22}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person01:
			fmt.Printf("list[%d] is a Person and its value is [%s, %d]\n", index, value.name, value.age)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}

}
