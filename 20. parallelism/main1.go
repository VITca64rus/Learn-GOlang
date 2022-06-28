package main

// import "fmt"

func task2(c chan string, s string) {
	for i:=0; i < 5; i++ {
		c <- s+" "
	}
}

func main() {
	c := make(chan string)
	go task2(c, "cat")
	print(<- c)
	print(<- c)
}