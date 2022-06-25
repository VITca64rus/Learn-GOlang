package main

import "fmt"

func main() {
	var a int
	var s string
	fmt.Scan(&a)
	if (a == 1) {
		s = "korova"
	} else if (a >= 2 && a <= 4) {
		s = "korovy"
	} else if (a >= 21 && (a%10==2 || a%10==3 || a%10==4)) {
		s = "korovy"
	} else if (a>=21 && a%10==1) {
		s = "korova"
	} else {
		s = "korov"
	}
	fmt.Printf("%d %s", a, s)
}