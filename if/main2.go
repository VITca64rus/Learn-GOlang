package main

import "fmt"

func main() {
	var n, a, b, c int

	fmt.Scan(&n)
	a = n % 10
	b = (n / 10) % 10
	c = n / 100

	if (a != b) && (a != c) && (b != c) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}