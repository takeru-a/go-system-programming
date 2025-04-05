package main
import (
	"compress/gzip"
	"os"
	"io"
)

func main() {
	file, err := os.Create("test.gz")
	if err != nil {
		panic(err)
	}

	writer := gzip.NewWriter(file)
	writer.Header.Name = "test"
	io.WriteString(writer, "gzip!!\n")
	writer.Close()
}

