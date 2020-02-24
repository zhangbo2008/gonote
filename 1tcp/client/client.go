package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	println("3333333333")
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	println("4444")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	println(3434343434343)
	fmt.Println(status,"内容是")    // status 就是内容
}
