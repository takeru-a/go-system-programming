package main
import (
	"bufio"
	"os"
)

func main() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("Hello!!\n")
	buffer.Flush()

	buffersize := bufio.NewWriterSize(os.Stdout, 10)
	buffersize.WriteString("01234567891")
}
