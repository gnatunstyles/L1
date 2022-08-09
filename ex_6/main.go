package main

import (
	"math/rand"
	"sync"
)

func main() {
	//при помощи сигнала канала выхода
	// data := make(chan int)
	// quit := make(chan struct{})
	// go func() {
	// 	for {
	// 		select {
	// 		case data <- mock():
	// 			fmt.Println("Sending")
	// 		case <-quit:
	// 			close(data)
	// 			return
	// 		}
	// 		time.Sleep(time.Millisecond * 100)
	// 	}
	// }()

	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	quit <- struct{}{}
	// }()

	// for i := range data {
	// 	fmt.Println(i)
	// }

	//при помощи waitgroup
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)
	go func() {
		for {
			foo, ok := <-ch
			if !ok {
				println("done")
				wg.Done()
				return
			}
			println(foo)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	wg.Wait()

}

func mock() int {
	return rand.Intn(100)
}
