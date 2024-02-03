package main

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/edsrzf/mmap-go"
)

func main() {
	test := []byte("aaabbbbxxxxxbealrkgjae:origjae")
	path := path.Join(os.TempDir(), "test")
	err := os.WriteFile(path, test, 0644)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	m, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer m.Unmap()
	m[9] = 'X'
	m.Flush()

	modifiedFile, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("original: %s\n", test)
	fmt.Printf("mmap: %s\n", m)
	fmt.Printf("file: %s\n", modifiedFile)
}
