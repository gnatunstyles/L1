package main

import (
	"fmt"
	"time"
)

func main() {
	var n time.Duration
	fmt.Scanln(&n)
	start := time.Now()
	fmt.Printf("Sleep for %d seconds\n", n)
	<-time.After(time.Second * n)
	fmt.Println("Program fineshed with time: ", time.Since(start))
}
