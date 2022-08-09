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
	//формирование массива из заготовленной строки
	arr := strings.Split(line[:len(line)-1], "")
	//формирование новой строки посредством чтения массива с другой стороны
	reversed := ""
	for i := len(arr) - 1; i >= 0; i-- {
		reversed += arr[i]
	}
	fmt.Println(reversed)
}
