package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	// 实例化一个socket
	listener, _ := net.Listen("tcp", ":10010")
	// 注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 启动服务
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go rpc.ServeConn(conn)
	}

}
