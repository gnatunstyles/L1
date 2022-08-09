package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b uint64
	var opt string
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Select option(+,-,/,*):")

	_, err = fmt.Scanln(&opt)
	switch opt {
	case "+":
		fmt.Println(a + b)
		return
	case "-":
		fmt.Println(a - b)
		return
	case "/":
		fmt.Println(a / b)
		return
	case "*":
		fmt.Println(a * b)
		return
	default:
		fmt.Println("ERROR. WRONG INPUT.")
		return
	}
}
