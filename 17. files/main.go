package main

import (
	"os"
	"strconv"
	"bufio"
	"io"
)

func main() {
	var res int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		q, err := strconv.Atoi(s)
		if err != nil && err != io.EOF {
			break
		}
		res += q
	}
	r := strconv.Itoa(res)
	io.WriteString(os.Stdout,r)
}