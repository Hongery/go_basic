package main

import (
	"studygo/day01/54niaobao_solve/proto"

	"net"
	"fmt"
)

// socket_stick/client/main.go

func main() {
	socket, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer socket.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		b,err:=proto.Encode(msg) //编码
		if err !=nil{
			fmt.Println("encode failed,err :",err)
			return
		}
		socket.Write(b)
		//socket.Write([]byte(msg))
	}
}