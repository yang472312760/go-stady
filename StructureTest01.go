package main

import (
	"fmt"
)

type Student struct {
	id int

	name string

	sex byte

	age int

	addr string
}

func TestA() {
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

func TestB() {

	//===============结构体变量为指针变量
	//3、先分配空间，再赋值
	s3 := new(Student)
	s3.id = 3
	s3.name = "xxx"
	fmt.Println(s3) //&{3 xxx 0 0 }

	//4、普通变量和指针变量类型打印
	var s4 Student = Student{4, "yyy", 'm', 18, "sz"}
	fmt.Printf("s4 = %v, &s4 = %v\n", s4, &s4) //s4 = {4 yyy 109 18 sz}, &s4 = &{4 yyy 109 18 sz}

	var p *Student = &s4
	//p.成员 和(*p).成员 操作是等价的
	p.id = 5
	(*p).name = "zzz"
	fmt.Println(p, *p, s4) //&{5 zzz 109 18 sz} {5 zzz 109 18 sz} {5 zzz 109 18 sz}

}

func TestC() {

	s1 := Student{1, "mike", 'm', 18, "sz"}
	s2 := Student{1, "mike", 'm', 18, "sz"}

	fmt.Println("s1 == s2", s1 == s2) //s1 == s2 true
	fmt.Println("s1 != s2", s1 != s2) //s1 != s2 false

}

func printStudentValue(tmp Student) {
	tmp.id = 250
	//printStudentValue tmp =  {250 mike 109 18 sz}
	fmt.Println("printStudentValue tmp = ", tmp)
}

func printStudentPointer(p *Student) {
	p.id = 250
	//printStudentPointer p =  &{250 mike 109 18 sz}
	fmt.Println("printStudentPointer p = ", p)
}

func main() {

	var s Student = Student{1, "mike", 'm', 18, "sz"}

	printStudentPointer(&s)     //引用(地址)传递，形参的修改会影响到实参
	fmt.Println("main s = ", s) //main s =  {250 mike 109 18 sz}

}
