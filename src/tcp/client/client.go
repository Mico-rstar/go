package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10010")
	if err != nil {
		log.Fatal(err)
	}
	data := []byte("hello, world")
	n, err := conn.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("send %d bytes\n", n)
	recvData := make([]byte, 1024)
	conn.Read(recvData)
	fmt.Println("Send to server: hello, world", )
	fmt.Println("Recv from server: ", string(recvData))

}
