package main

import "fmt"

func main() {
	var x int
	m := make(map[int]int)

	for i:=0; i < 10; i++ {
		fmt.Scan(&x)
		if m[x] != 0 {
			fmt.Printf("%d", m[x])
		} else {
			m[x] = work(x)
			fmt.Printf("%d", m[x])
		}
	}
}