package net

import (
	"net"
	"zinx/ziface"

)

type Connection struct {
	//当前连接的原生套接子
	Conn *net.TCPConn

	//链接ID
	ConnId uint32

	//当前的链接状态
	isClosed bool

	//当前链接锁绑定的业务处理方法
	handleAPI ziface.HandleFunc
}
/*

 */

func NewConnection(conn *net.TCPConn,connIP uint32,callback_api ziface.HandleFunc)ziface.IConnection  {
	c := &Connection{
		Conn:conn,
		ConnId:connIP,
		isClosed:false,
		handleAPI:callback_api,
	}
	return c
}

//启动链接
func (c *Connection) Start()  {


}
//停止链接
func (c *Connection)Stop()  {

}
//获取链接
func (c *Connection)GetConnID() uint32 {
	return 0

}
//获取conn的原生socket套接子
func (c *Connection)GetTCPConnection() *net.TCPConn  {
	return nil

}
//获取远程客户端的ip地址
func (c *Connection) GetRemoteAddr() net.Addr{
	return nil
}
//发送数据给对方客户端
func (c *Connection) Send(data []byte) error{
	return nil
}