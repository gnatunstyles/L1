package main

import (
	"fmt"
	"sync"
)

//структура счетчика
type Counter struct {
	val int
	m   sync.RWMutex
}

//конструктор
func NewCounter() *Counter {
	return &Counter{
		val: 0,
	}
}

//функция вызова инкремента
func (c *Counter) Increment() {
	c.m.Lock()
	defer c.m.Unlock()
	c.val++
}

//Функция конкурентного увеличения
func (c *Counter) parallelCount(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 150; i++ {
		c.Increment()
	}
}

func main() {
	var wg sync.WaitGroup
	//создание счетчика и вызов 3 параллельных горутин
	c := NewCounter()

	wg.Add(3)

	for i := 0; i < 3; i++ {
		go c.parallelCount(&wg)
	}

	wg.Wait()
	fmt.Println(c.val)
}
