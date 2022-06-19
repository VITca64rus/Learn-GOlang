/* ЗАДАНИЕ:
 * Напишите функцию sumInt, получающую переменное число аргументов типа int,
 * и возвращающую количество переданных аргументов и их сумму.
 */
 package main
 
 import "fmt"

func sumInt (a ...int) (int, int){
	var sum int = 0
	var count int = 0
	for _, i := range a {
		sum += i
		count += 1
	}
	return count, sum
}

func main() {
	a, b := sumInt(2, 4, 4)
	fmt.Println(a, b)
}