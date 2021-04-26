package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("./1.txt") //新建文件

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	for i := 0; i < 5; i++ {
		outstr := fmt.Sprintf("%s:%d\n", "Hello go", i)
		file.WriteString(outstr)
		file.Write([]byte("abcd\n"))
	}
}
