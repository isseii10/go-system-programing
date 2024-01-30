package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter")
}

func main() {
	fmt.Println("Hello, world!")

	// file, err := os.Create("test.txt")
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("aaaaaaaaaaaaaa\n"))
	defer file.Close()

	os.Stdout.Write([]byte("stdout\n"))
	os.Stderr.Write([]byte("これはエラー\n"))

	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer"))
	fmt.Println(buffer.String())

	var builder strings.Builder
	builder.Write([]byte("strings.Builder"))
	fmt.Println(builder.String())

	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ecample.com\r\n\r\n")
	io.Copy(os.Stdout, conn)

	http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// f, err := os.Create("multiwriter.txt")
	f, err := os.Open("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := io.MultiWriter(f, os.Stdout) // fileとstdoutどちらにも書き込む
	io.WriteString(writer, "io.MultiWriter example\n")
	_, err = io.CopyN(os.Stdout, file, 3)
	if err != nil {
		panic(err)
	}
}
