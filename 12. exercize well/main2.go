package main

import "fmt"

func main() {
	var s string

	fmt.Scan(&s)
	r := []rune(s)
	max := r[0]
	for _, c := range(r) {
		if c > max {
			max = c
		}
	}
	fmt.Println(string(max))
}