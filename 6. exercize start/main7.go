package  main

import "fmt"

func main(){
	var a, n,res int
	fmt.Scan(&a)
	for i:=0; i<a; i++ {
		fmt.Scan(&n)
		if n == 0 {
			res += 1
		}
	}

	fmt.Println(res)
}