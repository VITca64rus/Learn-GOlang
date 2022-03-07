package main

import "fmt"

func main() {
	var a []int
	var len, x int
	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Scan(&x)
		a = append(a, x)
	}
	fmt.Print(a[3])
}
