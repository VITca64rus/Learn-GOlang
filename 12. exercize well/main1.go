package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string

	fmt.Scan(&s)
	len := utf8.RuneCountInString(s)
	for _,c := range(s) {
		if (len > 1) {
			fmt.Printf("%s*", string(c))
		} else {
			fmt.Printf("%s", string(c))
		}
		len -= 1
	}
}