package main

import (
	"fmt"
	"os"
	"encoding/json"
)

// インターフェースの定義
type Talker interface {
	Talk()
}

// 構造体
type Greter struct {
	name string
}

type Greter2 struct {
	name string
	age int
}

type Greter3 struct {
	name string
}

// メソッド
func (g Greter) Talk() {
	fmt.Printf("Hello, myname is %s\n", g.name)
}

func (g2 Greter2) Talk() {
	fmt.Fprintf(os.Stdout,"Hello!!, myname is %s age is %d\n", g2.name, g2.age)
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("","	")
	encoder.Encode(map[string]string{
		"hello":"world!!",
	})
}

func (g3 Greter3) Printname() {
	fmt.Println(g3.name)
}

// エントリーポイント
func main() {
	// インターフェースの定義を満たす構造体のポインタのみ代入できる
	// Talk()メソッドが定義されているインターフェース
	var talker Talker = &Greter{"Mike"}
	talker.Talk()

	var talker2 Talker = &Greter2{"June", 19}
	talker2.Talk()

	// var talker3 Talker = &Greter3{"Error"}
}
