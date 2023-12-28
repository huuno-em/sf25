package main

import (
	"fmt"
	"log"
)

func main() {
	var data string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %s\n", data)
}
