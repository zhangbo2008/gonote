package main

/* 这个client 和server就是服务器简单例子.*/

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()   //等待clent发送请求. 发送过来了才会accept成功,否则堵塞
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	var buf = "wzzzz\n"
	n, err := conn.Write([]byte(buf))   // Write就是返回给client的
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	fmt.Println(buf)
	conn.Close()
}
