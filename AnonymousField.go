package main

import (
	"fmt"
)

type Persion struct {
	name string

	sex byte

	age int
}

type Student03 struct {
	Persion // 匿名字段，那么默认Student就包含了Person的所有字段

	id int

	addr string

	name string
}

func main() {
	var s Student03 //变量声明

	//给Student的name，还是给Person赋值？
	s.name = "mike"

	//{Person:{name: sex:0 age:0} id:0 addr: name:mike}
	fmt.Printf("%+v\n", s)

	//默认只会给最外层的成员赋值
	//给匿名同名成员赋值，需要显示调用
	s.Persion.name = "yoyo"
	//Person:{name:yoyo sex:0 age:0} id:0 addr: name:mike}
	fmt.Printf("%+v\n", s)

}
