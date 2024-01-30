package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("png/sample.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}

	// 以下binary.Read()の実験
	bNum := []byte{0x0, 0x0, 0x27, 0x10, 0x0, 0x0, 0x0, 0x11}
	r1 := bytes.NewReader(bNum)
	var num32 int32
	for {
		err = binary.Read(r1, binary.BigEndian, &num32) // int32なので4byteずつ読み込む
		if err == io.EOF {
			break
		}
		fmt.Printf("num: %d\n", num32)
	}
	r2 := bytes.NewReader(bNum)
	var num8 int8
	for {
		err = binary.Read(r2, binary.BigEndian, &num8) // int8なので1byteずつ読み込む
		if err == io.EOF {
			break
		}
		fmt.Printf("num: %d\n", num8)
	}
}

func readChunks(f *os.File) []io.Reader {
	var chunks []io.Reader
	f.Seek(8, 0) // 8byteとばす
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(f, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(f, offset, int64(length)+12))
		offset, _ = f.Seek(int64(length+8), 1)
	}
	return chunks
}

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk %s (%d bytes)\n", buffer, length)
}
