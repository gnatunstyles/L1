package main

import (
	"fmt"
	"time"
)

func main() {
	var n time.Duration
	//ввод необходимого значения с клавиатуры
	fmt.Scanln(&n)
	fmt.Printf("Sleep for %d seconds\n", n)
	start := time.Now()
	//вызов функции newtimer, которая остановит выполнение программы на n секунд
	<-time.NewTimer(n * time.Second).C
	fmt.Println("Program finished with time: ", time.Since(start))
}
