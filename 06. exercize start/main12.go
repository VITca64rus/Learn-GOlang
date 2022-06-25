package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Scan(&a)
	for i:=0; i <= a; i++ {
		if math.Log2(float64(i)) == float64(int(math.Log2(float64(i)))) {
			fmt.Printf("%d ", i)
		}
	}
}