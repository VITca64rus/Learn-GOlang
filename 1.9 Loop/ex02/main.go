package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	var res = 0
	for ; a <= b; a++ {
		res += a
	}
	fmt.Print(res)
}
