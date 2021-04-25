package main

import "fmt"

type Element interface {
}

type Person struct {
	name string
	age  int
}

func main() {

	list := make([]Element, 3)

	list[0] = 1
	list[1] = "hello"
	list[2] = Person{name: "yang", age: 14}

	for i, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", i, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", i, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is [%s, %d]\n", i, value.name, value.age)
		} else {
			fmt.Printf("list[%d] is of a different type\n", i)
		}
	}

}
