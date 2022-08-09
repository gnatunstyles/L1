package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//функция воркера
func worker(wg *sync.WaitGroup, id int, jobs <-chan int) {
	defer wg.Done()
	for i := range jobs {
		fmt.Printf("worker %d printed: %d\n", id, i)
	}
	fmt.Printf("worker %d finished\n", id)
}

func main() {
	fmt.Println("Enter the number of workers:")
	//количество воркеров
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(num)

	jobs := make(chan int, 1)
	//конкурентный запуск воркеров
	for i := 0; i < num; i++ {
		go worker(wg, i, jobs)
	}
	//определение сигнала остановки программы
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-c:
			close(jobs)
			time.Sleep(time.Millisecond * 10)
			fmt.Println("All channels stopped working")
			return
		default:
			jobs <- rand.Intn(100)
		}
		time.Sleep(time.Millisecond * 10)
	}

}
