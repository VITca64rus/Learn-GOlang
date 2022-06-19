package main

import "fmt"

func minimumFromFour() int {
    var a [4]int  
	fmt.Scan(&a[0])
	fmt.Scan(&a[1])
	fmt.Scan(&a[2])
	fmt.Scan(&a[3])

	min := a[0]
	for i:=1; i<4; i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}

func main(){
	fmt.Println(minimumFromFour())
}
