package main

import "fmt"
import "strconv"

func main() {
	var dec int64
	fmt.Scan(&dec)
	v := strconv.FormatInt(dec, 2)
	fmt.Println(v)
}