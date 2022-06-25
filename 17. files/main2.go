package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var i int
	file, _ := os.Open("test.txt")
	defer file.Close()
	rd := bufio.NewReader(file)
	for true {
		s, _ := rd.ReadString(';')
		if s == "0;" {
			break
		}
		i += 1
	}
	fmt.Println(i+1)
}