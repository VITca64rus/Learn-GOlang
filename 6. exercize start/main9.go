package main

import "fmt"

func main() {
	var a, res int
	fmt.Scan(&a)
	res = a
	for res / 10 != 0 {
		res = 0
		for a % 10 != 0 {
			res += (a % 10)
			a = a / 10
		}
		a = res
	}
	fmt.Println(res)
}