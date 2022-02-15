package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)
	a = a * 2
	a += 100
	fmt.Print(a)
}
