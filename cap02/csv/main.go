package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	csv := csv.NewWriter(file)
	csv.Write([]string{"hello", "world"})
	csv.Flush()
	file.Close()
}
