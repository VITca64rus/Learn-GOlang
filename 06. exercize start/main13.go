package main

import "fmt"

func fibi(n int) int {
    var a []int
	a = append(a, 1)
	a = append(a, 1)
	var i int = 2
    for a[i-1] != n {
        a = append(a, a[i - 1] + a[i-2])
		i += 1
		if a[i-1] > n {
			return (-1)
		}
    }
    return i
}

func main(){
	var a int
	fmt.Scan(&a)
	fmt.Println(fibi(a))
}

