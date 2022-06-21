package main

import "fmt"

func main() {
	var s1, s2 string
	var i, j int

	fmt.Scan(&s1, &s2)
	i = 0
	for i < len(s1) {
		j = 0
		for j < len(s2) {
			if s1[i] == s2[j] {
				fmt.Print(string(s1[i]), " ")
				break
			}
			j++
		}
		i++
	}
}