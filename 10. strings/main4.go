package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	for _, c := range(s) {
		if strings.Count(s, string(c)) <= 1 {
			fmt.Printf("%s", string(c))
		}
	}
}