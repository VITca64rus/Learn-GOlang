package main

import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/csv"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	buf, _ := os.Open(path)
	r := csv.NewReader(buf)
	data, err := r.ReadAll()
	// var res string
	i := 0
	for _, row := range data {
		i+=1
		row[0] = "S"
	}
	if i == 10 {
		print(path)
		for _, row := range data {
			fmt.Println(row)
		}
	}
	return (nil)
}

func main() {
	const root = "./task" // Файлы моей программы находятся в другой директории

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}