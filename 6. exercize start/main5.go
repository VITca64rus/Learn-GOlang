package  main

import "fmt"

func main(){
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a+b>c && a+c>b && c+b>a {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}