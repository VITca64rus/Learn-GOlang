package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, s string
	fmt.Scan(&x)
	fmt.Scan(&s)
	fmt.Println(strings.Index(x, s))
}