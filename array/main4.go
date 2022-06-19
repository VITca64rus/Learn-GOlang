package main

import "fmt"

func main() {
	var n, tmp,res int
	var arr []int

	res = 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		if tmp > 0 {
			res++
		} 
	}
	fmt.Println(res)
}