package main

import "fmt"

func main() {
	//инициализация переменных
	arr1 := []string{"a", "f", "d", "s", "c", "x", "z", "v"}
	arr2 := []string{"a", "f", "d", "q", "w", "x", "h", "g"}
	exist := make(map[string]bool)
	x := []string{}

	//итерация по первому массиву,добавление в map всех элементов
	for _, elt := range arr1 {
		exist[elt] = true
	}
	//итерация по 2 массиву, если в map элемент присутствует - то добавляем его в итоговый список пересечений
	for _, elt := range arr2 {
		if exist[elt] {
			x = append(x, elt)
		}
	}
	fmt.Println(x)
}
