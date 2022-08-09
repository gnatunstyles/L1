package main

import (
	"fmt"
	"log"
)

func main() {
	n := 0
	s := []string{"q", "w", "e", "r", "t", "y", "u", "i"}
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatal(err)
	}
	s = append(s[:n], s[n+1:]...)
	fmt.Println(s)
}
