package file

import (
	"os"
	"fmt"
)


func File_fmt(name string) {
	file, err := os.Create("test2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(file,"Hello!%s %d %f",name, 123, 0.11)
	file.Close()
}
