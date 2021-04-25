package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"-"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:",string"`
	Price    float64  `json:"price,omitempty"`
}

func main() {

	t1 := IT{"itcast", []string{"go", "c++", "python", "test"}, true, 666.666}

	b, err := json.Marshal(t1)

	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
