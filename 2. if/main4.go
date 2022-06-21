package main

import "fmt"

func main() {
	var a, sum1, sum2 int
	fmt.Scan(&a)
	sum1 = (a % 10) + ((a / 10) % 10) + ((a / 100) % 10)
	a = a / 1000
	sum2 = (a % 10) + ((a / 10) % 10) + ((a / 100) % 10)
	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}