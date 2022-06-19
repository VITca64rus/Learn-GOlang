package main

import "fmt"

func main() {
	var workArray [10]uint8
	var j, l, tmp uint8

	for i:=0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}
	for i:=0; i < 3; i++ {
		fmt.Scan(&j)
		tmp = workArray[j]
		fmt.Scan(&l)
		workArray[j] = workArray[l]
		workArray[l] = tmp
	}
	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}
}