package main

import (
	"fmt"
)

func main() {
	scala()
	pointer()
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
