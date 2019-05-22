package main

import (
	"zinx/net"
	"fmt"
)

func main() {
	//创建一个zinx server对象
	s := net.NewServer("zinx v0.1版本")

	//让server对象 启动服务
	s.Serve()
	fmt.Println("server启动")
}
