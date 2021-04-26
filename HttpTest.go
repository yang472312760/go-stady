package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	accept, err := listen.Accept()

	if err != nil {
		log.Println(err)
	}

	defer accept.Close()

	ipAddr := accept.RemoteAddr().String()

	fmt.Println(ipAddr, "连接成功")

	buf := make([]byte, 4096)

	n, err := accept.Read(buf)

	if err != nil {
		fmt.Println(err)
		return
	}

	//切片截取，只截取有效数据
	result := buf[:n]
	fmt.Printf("接收到数据来自[%s]==>:\n%s\n", ipAddr, string(result))

}
