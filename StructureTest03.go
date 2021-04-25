package main

import "fmt"

func main() {
	//===============结构体变量为普通变量
	//1、打印成员
	var s1 Student = Student{1, "mike", 'm', 18, "sz"}
	//结果：id = 1, name = mike, sex = m, age = 18, addr = sz
	fmt.Printf("id = %d, name = %s, sex = %c, age = %d, addr = %s\n", s1.id, s1.name, s1.sex, s1.age, s1.addr)

	//2、成员变量赋值
	var s2 Student
	s2.id = 2
	s2.name = "yoyo"
	s2.sex = 'f'
	s2.age = 16
	s2.addr = "guangzhou"
	fmt.Println(s2) //

}
