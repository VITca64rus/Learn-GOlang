package main

import "fmt"
import "strconv"

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	for _, c := range(r) {
		num, _ := strconv.Atoi(string(c))
		fmt.Printf("%d", num * num)
	}
}