package main

import (
	"net/http"
	"io"
	"compress/gzip"
	"encoding/json"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "Hello, World!!")
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string {
		"Hello": "World!!",
	}

	// gzip圧縮
	// 標準出力
	gzip := gzip.NewWriter(w)
	defer gzip.Close()
	writer := io.MultiWriter(gzip, os.Stdout)
	// json
	jsonec := json.NewEncoder(writer)
	jsonec.SetIndent("", "	")
	jsonec.Encode(source)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
