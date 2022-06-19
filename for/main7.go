package main

import "fmt"

func main() {
	var x, p, y, i int
	i = 0
	fmt.Scan(&x, &p, &y)
	for x < y {
		x += int(x * p / 100)
		i++
	}
	fmt.Println(i)
}