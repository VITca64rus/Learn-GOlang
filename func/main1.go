package main

import "fmt"

func vote(x int, y int, z int) int {
    var arr [3]int
	var count_zero int
	arr[0] = x
	arr[1] = y
	arr[2] = z

	count_zero = 0
	for i:=0; i < 3; i++ {
		if arr[i] == 0 {
			count_zero += 1
		}
	}
	if count_zero >= 2 {
		return (0)
	}
	return 1
}

func main() {
	fmt.Println(vote(0, 1, 1))
}