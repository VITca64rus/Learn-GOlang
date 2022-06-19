package main

import "fmt"

func main() {
	var n, a, res int
	res = 0
	fmt.Scan(&n)
	for i:=0; i < n; i++ {
		fmt.Scan(&a)
		if a > 9 && a < 100 && a % 8 == 0 {
			res += a
		}
	}
	fmt.Println(res)
}