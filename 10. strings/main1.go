package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string
	var j int
	fmt.Scan(&s)
	r := []rune(s)
	ns := []rune(s)
	len := utf8.RuneCountInString(s)
	j = 0
	for i := len - 1; i >= 0; i-- {
		ns[j] = r[i]
		j += 1
	}
	if string(r) == string(ns){
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}