package main

import (
	"fmt"
	"io"
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
	for {
		_, err = conn.Read(recvData)
		if err == io.EOF {	
			break
		}
	}
	fmt.Println("Send to server: hello, world", )
	fmt.Println("Recv from server: ", string(recvData))

}
