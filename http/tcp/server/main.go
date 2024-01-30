package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("start http server at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSettion(conn)
	}
}

func isGZipAcceptable(req *http.Request) bool {
	return strings.Index(
		strings.Join(req.Header["Accept-Encoding"], ","), "gzip") != -1
}

func processSettion(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("accept %v\n", conn.RemoteAddr())
	fmt.Println("==========================================")
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
		fmt.Println("==========================================")

		response := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}
		if isGZipAcceptable(request) {
			content := "hello world (gzipped)\n"
			var buf bytes.Buffer
			writer := gzip.NewWriter(&buf)
			writer.Write([]byte(content))
			writer.Close()
			response.Body = io.NopCloser(&buf)
			response.Header.Set("Content-Encoding", "gzip")
			response.ContentLength = int64(buf.Len())
		} else {
			content := "hello world\n"
			response.Body = io.NopCloser(strings.NewReader(content))
			response.ContentLength = int64(len(content))
		}
		response.Write(conn)

	}
}
