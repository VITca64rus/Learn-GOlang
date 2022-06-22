package main

import "fmt"
import "strconv"

func adding(a string, b string) int64 {
	var str1 string = ""
	var str2 string = ""
	r1 := []rune(a)
	r2 := []rune(b)
	for _, c := range(r1) {
		if c >= '0' && c <= '9'{
			str1 += string(c)
		}
	}
	for _, c := range(r2) {
		if c >= '0' && c <= '9'{
			str2 += string(c)
		}
	}
	fmt.Println(str1, str2)
	x, _ := strconv.Atoi(str1)
	x1, _ := strconv.Atoi(str2)
	return (int64(x + x1))
}

func main() {
	adding("#@cd5", "sfv4vf")
}