package main

import (
	"fmt"
	"time"
	"strings"
	"bufio"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	s, _ := rd.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

//13.03.2018 14:00:15,12.03.2018 14:00:15
	t := strings.Split(s, ",")
	t1, _ := time.Parse("02.01.2006 15:04:05", t[0])
	t2, _ := time.Parse("02.01.2006 15:04:05", t[1])
	if !t1.Before(t2) {
		res:= time.Since(t2).Round(time.Second) - time.Since(t1).Round(time.Second)
		fmt.Println(res)
	} else {
		res := time.Since(t1).Round(time.Second) - time.Since(t2).Round(time.Second)
		fmt.Println(res)
	}
}