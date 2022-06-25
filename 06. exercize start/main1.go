package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := n/100 + n%10 + n/10 % 10
	fmt.Println(res)
}