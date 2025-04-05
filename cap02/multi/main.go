package main

import (
	"io"
	"os"
)

func main() {
	// ファイル作成
	file, err := os.Create("multi.txt")
	if err != nil {
		panic(err)
	}

	// 複数のio.Writerを統合（ファイル出力、標準出力）
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "MultiWriter!!!")
}