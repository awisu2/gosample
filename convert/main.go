package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	scala()
	pointer()
	convert()
	stringToReader("abc")
}

func scala() {
	var i int
	i = 1

	type myInt int
	mi := myInt(i)
	mi = 2

	//Error:型を合わせずに渡すことはできない
	//i = mi
	i = int(mi)

	fmt.Println(i, mi)
}

func pointer() {
	fmt.Println("========== pointer ==========")
	type s struct{}
	type S s

	var sPtr *s
	sPtr = &s{}

	// ポインタをコンバートするときは型を()で囲む
	var SPtr *S
	SPtr = (*S)(sPtr)

	fmt.Println(sPtr, SPtr)
}

func convert() {
	fmt.Println("========== convert ==========")
	var i int
	var s string

	// int -> string
	i = 9999
	s = strconv.Itoa(i)

	// int:9999 => string:9999
	fmt.Printf("%T:%v => %T:%v\n", i, i, s, s)

	// string:9999 => int:9999
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T:%v => %T:%v\n", s, s, i, i)
}

// 文字列をIo.Readerに変換
func stringToReader(str string) {
	fmt.Println("========== stringToReader ==========")

	reader := strings.NewReader(str)
	fmt.Println(reader)
}
