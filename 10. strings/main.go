package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)
	len := utf8.RuneCountInString(text)
	if unicode.IsUpper(rs[0]) && rs[len - 2] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}