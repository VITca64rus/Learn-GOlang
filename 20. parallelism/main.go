package main

import "fmt"

func task(c chan int, N int) {
	c <- N+1
}

func main() {
	c := make(chan int)
	go task(c, 5)
	fmt.Println(<- c)
}