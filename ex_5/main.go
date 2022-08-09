package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//Функция, генерирующая случайные значения и отправляющая их в канал
func send(c chan int) {
	for {
		c <- rand.Intn(100)
	}
}

//Функция, читающая случайные значения из канала и выводящая их в stdout
func write(c chan int) {
	for {
		fmt.Println(<-c)
	}
}

func main() {
	//Ввод количества секунд работы программы
	fmt.Println("Enter the number of seconds:")
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatal(err)
	}

	//Инициализация канала
	c := make(chan int, 1)

	//Запуск функций отправки и чтения в отдельных горутинах
	go send(c)
	go write(c)

	//Время работы программы
	time.Sleep(time.Duration(num) * time.Second)
	//Закрытие канала
	close(c)
}
