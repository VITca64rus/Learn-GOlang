package main

import "fmt"

func main(){
	var s int
	fmt.Scan(&s)
	min := s / 60
	h := min / 60
	min = min % 60
	fmt.Println("It is", h, "hours", min, "minutes.")
}
