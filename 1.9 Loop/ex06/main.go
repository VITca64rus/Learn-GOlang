package main

import "fmt"

func main() {
	var n int
	for {
		fmt.Scan(&n)
		if n >= 10 && n <= 100 {
			fmt.Print(n, "\n")
		} else if n > 100 {
			break
		}
	}
}
