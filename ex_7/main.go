package main

import (
	"fmt"
	"sync"
)

//структура конкурентной MyMap
type MyMap struct {
	mx sync.RWMutex
	m  map[string]int
}

//конструктор MyMap
func NewMyMap() *MyMap {
	return &MyMap{
		m: make(map[string]int),
	}
}

//функция для загрузки данных в MyMap
func (c *MyMap) Load(key string, value int) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	c.m[key] = value
}

//функция для получения данных из map
func (c *MyMap) Get(key string) int {
	c.mx.RLock()
	item := c.m[key]
	c.mx.RUnlock()
	return item
}

func main() {
	myMap := NewMyMap()

	var wg sync.WaitGroup

	//конкурентно запускаем 3 горутины с записью данных в map

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			myMap.Load(fmt.Sprintf("%d", i), i)
		}
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			myMap.Load(fmt.Sprintf("%d", i), i)
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 20; i < 30; i++ {
			myMap.Load(fmt.Sprintf("%d", i), i)
		}
	}()

	wg.Wait()
	//выводим результат выполнения на экран
	fmt.Println(myMap)
}
