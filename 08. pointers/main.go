package main

import "fmt"

func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}

func main() {
	var a int = 5
	var b int = 2
	test(&a, &b)
}