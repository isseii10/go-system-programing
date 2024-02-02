package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("server is runnning at localhost:8888")
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer) // ReadFrom: connから読み取ってbufferにコピーする
		if err != nil {
			panic(err)
		}
		fmt.Printf("remote address %v: length(%v)", remoteAddress, length)
		_, err = conn.WriteTo([]byte("hello from server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}
