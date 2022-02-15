package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)
	a = a * a
	b = b * b
	a += b
	fmt.Print(a)
}
