package main

import "fmt"

func test(x1 *int, x2 *int) {
	var a int
	a = *x1
	*x1 = *x2
	*x2 = a
	fmt.Println(*x1, *x2)
}

func main() {
	var a int = 5
	var b int = 2
	test(&a, &b)
}