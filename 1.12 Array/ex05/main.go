package main

import "fmt"

func main() {
	var a []int
	var len, x int
	var count int
	count = 0
	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Scan(&x)
		a = append(a, x)
	}
	for i := 0; i < len; i++ {
		if a[i] > 0 {
			count++
		}
	}
	fmt.Print(count)
}
