package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Print(string(a[0]))
}
