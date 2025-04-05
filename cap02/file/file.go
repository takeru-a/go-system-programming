package file

import ( 
	"os"
)

func File() {
	// ファイル生成
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	
	file.Write([]byte("os.file example\n"))
	file.Close()
}