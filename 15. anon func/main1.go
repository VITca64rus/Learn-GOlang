package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(a uint) uint {
		var res string = ""
		s := strconv.FormatUint(uint64(a), 10)
		r := []rune(s)
		for _,c := range(r) {
			if c != '0' && c % 2 == 0 {
				res += string(c)
			}
		}
		rr, _ := strconv.ParseUint(res, 10, 64)
		if rr == 0{
			return (uint(100))
		} else {
			return (uint(rr))
		}
	}

	fmt.Println(fn(717171))
}