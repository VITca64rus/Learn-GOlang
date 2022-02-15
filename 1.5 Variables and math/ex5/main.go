package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	a /= 10
	a %= 10
	fmt.Print(a)
}
