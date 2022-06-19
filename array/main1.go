package main

import "fmt"

func main() {
	var n, tmp int
	var arr []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		arr = append(arr, tmp)
	}
	print(arr[3])
}