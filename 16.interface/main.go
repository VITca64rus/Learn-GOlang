package main

import "fmt"

func readTask() (value1, value2, operation interface{}) {

return 4.0, 2.0, "/" //тут играемся с параметрами

}



func printError(value interface{}) {
    fmt.Printf("value=%v: %T", value, value)
}

func main(){
	value1, value2, operation := readTask()
	var a, b float64

	if v, ok := value1.(float64); !ok {
		printError(value1)
		return
	} else {
		a = v
	}
	if v, ok := value2.(float64); !ok {
		printError(value2)
		return
	} else {
		b = v
	}
	if v, ok := operation.(string); !ok {
		fmt.Println("неизвестная операция")
		return
	} else {
		if v == "+" {
			fmt.Printf("%.4f", a + b)
		} else if v == "-" {
			fmt.Printf("%.4f", a - b)
		} else if v == "*" {
			fmt.Printf("%.4f", a * b)
		} else if v == "/" {
			fmt.Printf("%.4f", a / b)
		}
	}
}