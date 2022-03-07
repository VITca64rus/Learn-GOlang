package main

import "fmt"

func main() {
	var workArray [10]uint8
	var tmp uint8
	var i, j, l int

	for i = 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}
	for i = 0; i < 3; i++ {
		fmt.Scan(&j)
		fmt.Scan(&l)
		tmp = workArray[j]
		workArray[j] = workArray[l]
		workArray[l] = tmp
	}
	for i = 0; i < 10; i++ {
		fmt.Print(workArray[i])
		fmt.Print(" ")
	}
}
