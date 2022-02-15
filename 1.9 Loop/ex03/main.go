package main

import "fmt"

func main() {
	var n int
	var res = 0
	var num int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if (num%8 == 0) && (num/10 > 0) && (num/100 == 0) {
			res += num
		}
	}
	fmt.Print(res)
}
