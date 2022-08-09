package main

import (
	"fmt"
)

//Функция для определения типа интерфейса
func typeDetect(n interface{}) {
	switch v := n.(type) {
	case int:
		fmt.Println("int")

	case string:
		fmt.Println("string")

	case bool:
		fmt.Println("bool")

	case chan interface{}:
		fmt.Println("channel")

	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	//инициализация тестового массива
	var s []interface{}
	c := make(chan interface{})
	s = append(s, 11)
	s = append(s, "123")
	s = append(s, true)
	s = append(s, c)
	//итерация по тестовому массиву и вызов функции typeDetect для каждого из его элементов
	for _, elt := range s {
		typeDetect(elt)
	}
}
