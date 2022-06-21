package main

import (
	"fmt"
)

func main() {
	var s string
	var count_s, count_d int
	fmt.Scan(&s)
	rs := []rune(s)

	for _, c := range (rs) {
		if c >= 'a' && c <= 'z' {
			count_s += 1
		} else if c >= 'A' && c <= 'Z' {
			count_s += 1
		} else if c >= '0' && c <= '9' {
			count_d += 1
		} else {
			fmt.Println("Wrong password")
			return
		}
	}

	if count_d > 0 && count_s > 0 && count_d + count_s >= 5 {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}