package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	fmt.Scan(&s)
	t, _ := time.Parse(time.RFC3339, s)
	fmt.Println(t.Format(time.UnixDate))
}