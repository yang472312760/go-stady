package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	//创建、监听socket
	listen, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		log.Fatal(err) //log.Fatal()会产生panic
	}

	defer listen.Close()

	for {
		accept, err := listen.Accept() //阻塞等待客户端连接
		if err != nil {
			log.Println(err)
			continue
		}

		go dealConn(accept)
	}

}

func dealConn(conn net.Conn) {

	defer conn.Close() //此函数结束时，关闭连接套接字

	//conn.RemoteAddr().String()：连接客服端的网络地址
	ipAddr := conn.RemoteAddr().String()

	fmt.Println(ipAddr, "连接成功")

	buf := make([]byte, 1024) //缓冲区，用于接收客户端发送的数据

	for {
		//阻塞等待用户发送的数据
		read, err := conn.Read(buf) //n代码接收数据的长度
		if err != nil {
			fmt.Println(err)
			return
		}
		//切片截取，只截取有效数据
		bytes := buf[:read]
		fmt.Printf("接收到数据来自[%s]==>[%d]:%s\n", ipAddr, read, string(bytes))
		if "exit" == string(bytes) {
			fmt.Println(ipAddr, "退出连接")
			return
		}
		//把接收到的数据转换为大写，再给客户端发送
		conn.Write([]byte(strings.ToUpper(string(bytes))))
	}
}
