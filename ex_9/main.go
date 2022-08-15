package main

import (
	"fmt"
	"sync"
)

// square - считает квадраты чисел
func square(in chan int, out chan int) {
	for val := range in {
		square := val * val
		out <- square
	}
	close(out)
}

// прочесть данные с канала
func read(out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range out {
		fmt.Println(val)
	}
}

func main() {
	var wg sync.WaitGroup
	in := make(chan int, 10)
	out := make(chan int, 10)
	arr := [5]int{1, 2, 3, 4, 5}

	wg.Add(1)
	go square(in, out)
	go read(out, &wg)

	for _, val := range arr {
		in <- val
	}

	close(in)
	wg.Wait()
}
