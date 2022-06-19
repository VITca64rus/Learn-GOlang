package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	for i < 100 {
		if (i > 10) {
			fmt.Println(i)
		}
		fmt.Scan(&i)
	}
}