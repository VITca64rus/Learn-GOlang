package main

import "fmt" // пакет используется для проверки ответа, не удаляйте его

type Battery struct {
	zero string
	one string
}

func (b Battery) String() string {
	return fmt.Sprintf("[%s%s]", b.zero, b.one)
}

func main() {
	var s string
	var batteryForTest Battery
	fmt.Scan(&s)
	r := []rune(s)
	for _, c := range(r) {
		if c == '0' {
			batteryForTest.zero += string(' ')
		} else {
			batteryForTest.one += string('X')
		}
	}
}