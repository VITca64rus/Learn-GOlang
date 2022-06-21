package main

import "fmt"

func main() {
	var a, b, res int
	res = 0

	fmt.Scan(&a, &b)
    for i := a; i < b+1; i++ {
        res += i
    }
	fmt.Println(res)
}