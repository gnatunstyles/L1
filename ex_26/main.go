package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "qwertyuQiasdfghjkzxcvbnm"
	//создание массива из строни и преобразование элементов к нижнему регистру
	arr := strings.Split(strings.ToLower(str), "")
	//создание map для проверки
	m := make(map[string]bool)
	//итерация по массиву, если элемент с таким ключом уже есть в map - возвращаем false
	for i := range arr {
		if !m[arr[i]] {
			m[arr[i]] = true
		} else {
			fmt.Println("false")
			return
		}
	}
	//если после итерации не был возвращено false - возвращаем true, т.к. в массиве дубликатов обнаружено не было
	fmt.Println("true")
	return
}
