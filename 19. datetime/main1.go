package main

import (
	"fmt"
	"time"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	s, _ := rd.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t1, _ := time.Parse("2006-01-02 15:04:05", s)
	hour, _ := strconv.Atoi(s[11:13])
	if hour < 13 {
		fmt.Println(t1.Format("2006-01-02 15:04:05"))
	} else {
		t2 := t1.AddDate(0, 0, 1)
		fmt.Println(t2.Format("2006-01-02 15:04:05"))
	}
}