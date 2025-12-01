package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

// 统计字符数
func factory(str string) string {
	return strconv.Itoa(len(str))		
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10010")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", addr) 
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(conn.RemoteAddr().String())
		recv := make([]byte, 0, 1024)
		conn.Read(recv)
		sendData := []byte(factory(string(recv)))
		conn.Write(sendData)
		conn.Close()
	}
}