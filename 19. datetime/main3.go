package main

import (
	"fmt"
	"time"
	"strings"
	"bufio"
	"os"
)

const now = 1589570165

func main() {
	rd := bufio.NewReader(os.Stdin)
	s, _ := rd.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	//12 мин. 13 сек.
	s = strings.Replace(s, " мин. ", "m", -1)
	s = strings.Replace(s, " сек.", "s", -1)
	dur, _ := time.ParseDuration(s)
	t1 := time.Unix(now, 0)
	res := t1.Add(dur)
	fmt.Println(res.Format(time.UnixDate))
}