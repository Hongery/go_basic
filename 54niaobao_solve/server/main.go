package main

import (
	"net"
	"bufio"
	//"io"
	"fmt"
	"studygo/day01/54niaobao_solve/proto"
)

// socket_stick/server/main.go  黏包

func process(socket net.Conn) {
	defer socket.Close()
	reader := bufio.NewReader(socket)
	//var buf [1024]byte
	for {
		/*n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}*/
		//recvStr := string(buf[:n])
		recvStr,err :=proto.Decode(reader)
		if err !=nil{
			fmt.Println("decode failed,err:",err)
			return
		}
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		socket, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(socket)
	}
}