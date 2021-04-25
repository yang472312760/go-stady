package main

import "fmt"

type Humaner interface {
	SayHi()
}

type Personer interface {
	Humaner
	Sing(lyrics string)
}

type Student04 struct {
	name  string
	score float64
}

func (s *Student04) SayHi() {
	fmt.Printf("Student[%s, %f] say hi!!\n", s.name, s.score)
}

func (s *Student04) Sing(lyrics string) {
	fmt.Printf("Student sing[%s]!!\n", lyrics)
}

func main() {

	//var i1 Humaner = &Student{"mike", 88.88}
	//var i2 Personer = i1 //err

	//Personer为超集，Humaner为子集
	var i1 Personer = &Student04{"mike", 88.88}
	var i2 Humaner = i1
	i2.SayHi() //Student[mike, 88.880000] say hi!!
}
