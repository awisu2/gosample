package main

import (
	"fmt"
)

func main() {
	sample("a")           // a
	sample("b")           // b or c
	sample("c")           // b or c
	sample("fallthrough") // fallthrough fallthrough next
	sample("foo")         // default
}

func sample(s string) {

	r := ""

	// 文字列での分岐もOK
	switch s {
	case "a":
		r += "a"
	case "b", "c":
		r += "b or c"
	case "fallthrough":
		r += "fallthrough"
		fallthrough // case内の最後に記述する必要あり
	case "fallthrough next":
		r += " fallthrough next"
	default:
		r = "default"
	}

	fmt.Println(r)
}
