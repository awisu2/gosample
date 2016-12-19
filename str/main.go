package main

import (
	"fmt"
)

func main() {
	substr()
}

func substr() {
	s := "abc"

	// abc
	fmt.Println(s)

	// b
	fmt.Println(s[1:2])

	// bc
	fmt.Println(s[1:])

	// abc
	fmt.Println(s[:len(s)])
}
