package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:10010")
	if err != nil {
		panic(err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "bbcc", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println("收到服务端的恢复", reply)


}