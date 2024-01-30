package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("----------header----------\n")
	content := bytes.NewBufferString("text\n")
	footer := bytes.NewBufferString("----------footer----------\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)

	fmt.Println("===================")

	var buffer bytes.Buffer
	reader2 := bytes.NewBufferString("text2")
	teeReader := io.TeeReader(reader2, &buffer)
	_, _ = io.ReadAll(teeReader)
	fmt.Println(buffer.String())
}
