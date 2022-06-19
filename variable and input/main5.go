package main

import "fmt"

func main() {

  var degree, hour, minutes int
  fmt.Scan(&degree)

  hour = degree / 30
  minutes = (degree % 30) * 2
  fmt.Println("It is", hour, "hours", minutes, "minutes.")
}