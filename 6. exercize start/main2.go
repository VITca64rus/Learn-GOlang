package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(n%10)
	fmt.Print(n/10 % 10)
	fmt.Print(n/100)
}