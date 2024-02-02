package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp4", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("hello from client"))
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 1500)
	_, err = conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buffer))
}
