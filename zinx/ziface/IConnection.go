package ziface

import "net"

type IConnection interface {
	//启动连接
	Start()

	//停止连接
	Stop()

	//获取连接
	GetConnID() uint32

	//获取conn的原生socket套接字
	GetTCPConnection() *net.TCPConn

	//
	GetRemoteAddr() net.Addr

	//
	Send(data []byte) error
}
//业务处理方法 抽象定义
type HandleFunc func(*net.TCPConn,[]byte,int) error
