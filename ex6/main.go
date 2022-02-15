package main

import "fmt"

func main() {
	var degree int
	var hour int
	var minute int
	fmt.Scan(&degree)
	hour = degree / 30
	minute = 2 * (degree % 30)
	fmt.Print("It is ", hour, " hours ", minute, " minutes.")
}
