package main

import (
	"fmt"
	"time"
	"net"
)

func main(){
	fmt.Println("client start ---")

	time.Sleep( 1 *time.Second)

	conn,err := net.Dial("tcp","127.0.0.1:8999")
	if err!=nil{
		fmt.Println("client start err",err)
		return
	}
	for{
		_,err:=conn.Write([]byte("Hello zixn--"))
		if err!=nil{
			fmt.Println("write conn err",err)
			return
		}
		//du
		buf := make([]byte,512)
		cnt,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("server call back : %s,cnt = %d\n",buf,cnt)
		time.Sleep(1 *time.Second)
	}
	//time.Second(1 * time.Second)
}
