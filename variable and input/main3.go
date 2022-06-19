package main

import "fmt"

func main() {

  var a int
  fmt.Scan(&a) // считаем переменную 'a' с консоли

  a = a % 10
  fmt.Println(a)
}