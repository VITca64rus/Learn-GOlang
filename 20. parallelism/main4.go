package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1) // Увеличиваем счетчик горутин в группе
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			work()
		}(wg)
	}
	wg.Wait() // ожидаем завершения всех горутин в группе
	fmt.Println("Горутины завершили выполнение")
}

func work()  {
	fmt.Printf("Горутина начала выполнение \n")
	time.Sleep(2 * time.Second)
	fmt.Printf("Горутина завершила выполнение \n")
}