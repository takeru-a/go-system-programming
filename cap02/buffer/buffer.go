package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main () {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	buffer.WriteString("WriteString\n")
	// bufferはio.Writerインターフェースの実装である
	io.WriteString(&buffer, "io.WriteString\n")
	fmt.Println(buffer.String())

	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}