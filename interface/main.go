package main

import (
	"fmt"
)

type I1 interface {
	Calc(a int, b int) int
}

// 足し算
type Plus struct{}

func (self Plus) Calc(a int, b int) int {
	return a + b
}

// 掛け算
type Multipl struct{}

func (self Multipl) Calc(a int, b int) int {
	return a * b
}

// AnyValue
type Any interface{}

func main() {
	sameFunc(Plus{}, 2, 3)    // 5
	sameFunc(Multipl{}, 2, 3) // 6

	anyValue()

	intefaceMaps()
}

// 関数がそろっていれば処理を呼ぶことが可能
// ただし、プロパティ値にはアクセスできない
func sameFunc(in I1, a int, b int) {
	fmt.Println(in.Calc(a, b))
}

// どんな値でも受け付ける
func anyValue() {
	var i Any = 1
	var s Any = "s"
	var b Any = true
	var p Any = Plus{}

	fmt.Println(i) // 1
	fmt.Println(s) // s
	fmt.Println(b) // true
	fmt.Println(p) // {}

	// 1
	if v, ok := i.(int); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not int")
	}

	// not int
	if v, ok := s.(int); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not int")
	}

	// map[b:true p:{} i:1 s:s]
	m := map[string]Any{
		"i": i,
		"s": s,
		"b": b,
		"p": p,
	}
	fmt.Println(m)
}

func intefaceMaps() {
	fmt.Println("===== intefaceMaps =====")

	imap := map[string]I1{}
	imap["plus"] = Plus{}
	imap["multipl"] = Multipl{}

	for k, v := range imap {
		fmt.Println(k, v.Calc(2, 3))
	}
}
