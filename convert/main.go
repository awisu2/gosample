package main

import (
	"fmt"
	"strconv"
)

func main() {
	scala()
	pointer()
	convert()
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
	var i int
	var s string

	// int -> string
	i = 9999
	s = strconv.Itoa(i)

	// int:9999 => string:9999
	fmt.Printf("%T:%v => %T:%v", i, i, s, s)
}
