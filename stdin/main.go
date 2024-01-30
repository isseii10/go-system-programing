package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Printf("size=%v, %s\n", size, string(buffer))
	}
}
