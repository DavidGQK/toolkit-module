package main

import (
	"fmt"
	"github.com/davidgqk/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(3)
	fmt.Println("Random string:", s)
}
