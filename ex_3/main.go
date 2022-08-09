package main

import (
	"fmt"
	"sync"
)

func main() {
	//Инициализация массива
	arr := [5]int{2, 4, 6, 8, 10}
	sum := 0
	//Инициализация группы ожидания выполнения горутин
	var wg sync.WaitGroup

	//Итерация по массиву
	for i := range arr {
		//Добавление новой горутины в waitgroup
		wg.Add(1)
		//конкурентный запуск функции вывода
		go func(i int) {
			//отложенный вызов декрементации waitgroup
			defer wg.Done()
			//вывод квадрата значения массива по индексу в stdout
			sum += arr[i] * arr[i]
		}(i)
	}
	//ожидание выполнения waitgroup
	wg.Wait()

	fmt.Println(sum)
}
