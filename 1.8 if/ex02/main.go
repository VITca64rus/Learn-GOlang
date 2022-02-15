package main

import "fmt"

func main() {
	var a, x, y int

	fmt.Scan(&a)
	x = a % 10
	a /= 10
	y = a % 10
	a /= 10
	if (a != y) && (a != x) && (y != x) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
