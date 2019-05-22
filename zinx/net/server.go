package net

import (
	"zinx/ziface"
	"fmt"
	"net"
)

type Server struct {
	//服务器ip
	IPVersion string
	IP string
	Port int
	Name string
}

func NewServer(name string)ziface.IServer{
	s:= &Server{
		Name:name,
		IPVersion:"tcp4",
		IP:"0.0.0.0",
		Port:8999,

	}
	return s
}
//启动服务器
func (s *Server) Start() {
	fmt.Errorf("【start】Server Linstenner at IP:%s,Post:%d, is starting..\n",s.IP,s.Port)

	//1.创建套接字：得到一个tcp的addr
	addr,err:=net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))
	if err != nil{
		fmt.Println("resolve tcp addr error:",err)
		return
	}
	//2.监听服务器地址
	listenner,err:=net.ListenTCP(s.IPVersion,addr)
	if err !=nil {
		fmt.Println("listen",s.IPVersion,"err:",err)
		return
	}

	//3
	go func() {
		for{
			//
			conn,err:=listenner.Accept()
			if err!=nil{
				fmt.Println("Accept err",err)
				continue
			}

			go func() {
				for{
					buf:=make([]byte,521)
					cnt,err:= conn.Read(buf)
					if err!=nil{
						fmt.Println("recv buf err:",err)//EOF
						break
					}
					fmt.Printf("recv client buf%s,cnt = %d\n",buf,cnt)

					//回显功能(业务)
					_,err =conn.Write(buf[:cnt])
					if err!=nil{
						fmt.Println("write back buf err ",err)
						continue
					}
				}
			}()
		}
	}()
}
//停止服务器
func (s *Server) Stop() {

}
//运行服务器
func (s *Server )Serve() {
	//启动server的监听功能
	s.Start()

	//TODO 在一些其他的扩展
	//阻塞  告诉CPU不再需要处理的，节省cpu资源
	select {}
}