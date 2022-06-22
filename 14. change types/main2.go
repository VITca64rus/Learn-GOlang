package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, ",", ".", -1)
	arr := strings.Split(s, ";")
	n0, _ := strconv.ParseFloat(arr[0], 64);
	n1, _ := strconv.ParseFloat(arr[1], 64);
	fmt.Printf("%.4f", n0 / n1)
}