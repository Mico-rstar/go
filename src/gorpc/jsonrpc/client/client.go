package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:10010")
	if err != nil {
		panic(err)
	}
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "bbcc", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println("收到服务端的恢复", reply)


}