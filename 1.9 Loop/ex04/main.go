package main

import "fmt"

func main() {
	var i int
	var max int = 0
	var count int = 0
	for fmt.Scan(&i); i != 0; fmt.Scan(&i) {
		if i > max {
			max = i
			count = 1
		} else if i == max {
			count++
		}
	}
	fmt.Print(count)
}
