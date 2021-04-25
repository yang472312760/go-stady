package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 创建一个保存键值对的映射
	t1 := make(map[string]interface{})
	t1["company"] = "itcast"
	t1["subjects "] = []string{"Go", "C++", "Python", "Test"}
	t1["isok"] = true
	t1["price"] = 666.666

	b, err := json.Marshal(t1)
	//json.MarshalIndent(t1, "", "    ")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))

}
