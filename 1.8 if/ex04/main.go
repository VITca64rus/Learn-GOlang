package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	if a[0]+a[1]+a[2] == a[3]+a[4]+a[5] {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
