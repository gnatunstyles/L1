package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b uint64
	var opt string
	//ввод 2х числовых переменных
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		log.Fatal(err)
	}

	if (a > (1 << 20)) && (b > (1 << 20)) {
		//выбор операции и исполнение
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
}
