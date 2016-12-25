package main

import (
	"fmt"
	"github.com/russross/blackfriday"
)

func main() {
	output := blackfriday.MarkdownCommon([]byte("# h1\n* a\n* b"))
	fmt.Println(string(output))
}
