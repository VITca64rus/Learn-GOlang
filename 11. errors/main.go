package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)

	res, err := divide(x, y)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println("ошибка")
	}
}