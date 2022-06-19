package main

import "fmt"

func main() {
	var i, res, max int
	res = 0
	fmt.Scan(&i)
	max = i
	for i != 0 {
		if i == max {
			res += 1
		} else if i > max {
			max = i
			res = 1
		}
		fmt.Scan(&i)
	}
	fmt.Println(res)
}