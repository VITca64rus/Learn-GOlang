package main

import "fmt"

func main() {
	var n string
	var cc string
	fmt.Scan(&n)
	fmt.Scan(&cc)
	r := []rune(n)
	rr := []rune(cc)
	for _, c := range (r) {
		if c != rr[0] {
			fmt.Printf("%s", string(c))
		}
	}
}