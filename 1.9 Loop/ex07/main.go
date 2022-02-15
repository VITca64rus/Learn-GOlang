package main

import "fmt"

func main() {
	var x, p, y int
	var i = 0
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	for x < y {
		x += x * p / 100
		i++
	}
	fmt.Print(i)
}
