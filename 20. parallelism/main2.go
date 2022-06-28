package main

 import (
	"fmt"
 )

func removeDuplicates(s1 chan string, s2 chan string) {
	var a string
	for true {
		c, ok := <-s1
		if !ok {
			close(s2)
			break
		} else {
			if a != c {
				a = c
				s2 <- c
			}
		}
	}
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	
	go func() {
	   defer close(inputStream)
	
	   for _, r := range "112334456" {
		  inputStream <- string(r)
	   }
	}()
	
	for x := range outputStream {
	   fmt.Print(x)
	}
}