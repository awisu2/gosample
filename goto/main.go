package main

import (
	"fmt"
)

// 0
// 1
// 2
// jump G
func main() {

	i := 0
L:
	fmt.Println(i)

	i++
	if i < 3 {
		goto L
	}

	goto G

	// Error : goto G jumps over declaration of a
	// コメントを外すとエラー、gotoでメモリ確保は通り越せない
	// a := 0

	fmt.Println("no print")
	return
G:
	fmt.Println("jump G")
}
