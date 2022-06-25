package  main

import "fmt"

func main(){
	var a, n,res, min int
	fmt.Scan(&a)
	min = 10000
	for i:=0; i<a; i++ {
		fmt.Scan(&n)
		if n < min {
			min = n
			res = 1
		} else if (n == min) {
			res += 1
		}
	}

	fmt.Println(res)
}