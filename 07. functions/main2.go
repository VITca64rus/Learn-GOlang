package main

import "fmt"

func fibonacci(n int) int {
	var arr []int
	var i int = 0
    for i = 0; i <= n; i++ {
		if i == 0 || i == 1 {
			arr = append(arr, i)
		} else {
			arr = append(arr, arr[i - 1] + arr[i - 2])
		}
	}
	fmt.Println(arr[n])
	return(arr[n])
}

func main() {
	fibonacci(5)
}