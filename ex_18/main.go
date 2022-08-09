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

func main() {
	var wg sync.WaitGroup
	//создание счетчика и вызов 3 параллельных горутин
	c := NewCounter()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 150; i++ {
			c.Increment()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 150; i++ {
			c.Increment()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 150; i++ {
			c.Increment()
		}
	}()

	wg.Wait()
	fmt.Println(c.val)
}
