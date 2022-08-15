package main

import (
	"fmt"
	"math/rand"
)

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	//выбор опорного элемента
	pivot := rand.Int() % len(a)

	//делаем опорный элемент на последнее место для сравнения его с элементами массива

	a[pivot], a[right] = a[right], a[pivot]
	//итерация по массиву
	//если сравниваемый элемент меньше опорного - ставим его на место left, left итерируем
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	//ставим опорный элемент на место left
	a[left], a[right] = a[right], a[left]

	//рекурсивно вызываем функцию quicksort для массивов от и после опорного элемента
	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
func main() {
	fmt.Println(quicksort([]int{5, 6, 7, 2, 1, 0}))
	fmt.Println(rand.Int())
}
