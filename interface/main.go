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

func main() {
	sample(Plus{}, 2, 3)    // 5
	sample(Multipl{}, 2, 3) // 6
}

func sample(in I1, a int, b int) {
	fmt.Println("Calc : ", in.Calc(a, b))
}
