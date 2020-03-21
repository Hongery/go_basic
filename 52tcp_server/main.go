package main

import (
	"net"
	"fmt"
)

//tcp_服务端
//1. 监听端口
//2. 接收客户端请求建立链接
//3. 创建goroutine处理链接。

func processConn(conn net.Conn){
	defer conn.Close()
	//3.与客户端进行通信
	var tmp [128]byte
	for{
		n,err :=conn.Read(tmp[:])
		if err != nil{
			fmt.Println("read from conn failed,err:",err)
			return
	}
	fmt.Println(string(tmp[:n]))
}
}

func main(){
	//1.本地端口服务启动
	listener,err := net.Listen("tcp","127.0.0.1:20")
	if err != nil{
		fmt.Println("start tcp server on 127.0.0.1:20 failed,err :",err)
		return
	}
	//2.等待别人来跟我建立连接
	for{
		conn,err := listener.Accept()
		if err != nil{
			fmt.Println("accept failed ,err :",err)
			return
		}
		go processConn(conn)
	}

}
