package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//чтение строки из буфера
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//формирование массива
	arr := strings.Split(line[:len(line)-1], " ")

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf(arr[i] + " ")
	}
}
