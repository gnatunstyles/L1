package main

import (
	"fmt"
	"log"
)

func main() {
	n := 0
	s := []string{"q", "w", "e", "r", "t", "y", "u", "i"}
	//ввод индекса исключаемого элемента
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatal(err)
	}
	//прибавление 2х слайсов друг к другу без исключаемого элемента
	s = append(s[:n], s[n+1:]...)
	fmt.Println(s)
}
