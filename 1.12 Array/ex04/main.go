package main

import "fmt"

func main() {
	var len int
	var arr [100]int

	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < len; i += 2 {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
}
