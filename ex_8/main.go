package main

import (
	"fmt"
	"log"
)

// Меняет бит с 1 на 0 и наоборот
func change(n int64, pos uint) int64 {
	pos--
	val := n & (1 << pos)
	if val > 0 {
		n &= ^(1 << pos)
	} else {
		n |= (1 << pos)
	}
	return n
}

func main() {
	var n int64
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatal()
	}
	fmt.Printf("Number: %d\n in bytes: %08b\n", n, n)

	changedNum := change(n, 1)
	fmt.Printf("Number: %d\n in bytes:  %08b\n", changedNum, changedNum)
}
