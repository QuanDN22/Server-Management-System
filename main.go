package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, sem chan int) {
	defer wg.Done()
	fmt.Printf("Worker %d is starting\n", id)
	time.Sleep(2 * time.Second) // Giả lập công việc tốn thời gian
	fmt.Printf("Worker %d is done\n", id)
	<-sem // Giải phóng một slot
}

func main() {
	log.Println("config parsed...")

	const numWorkers = 20
	const maxConcurrentWorkers = 3 // Giới hạn số lượng goroutine đồng thời

	var wg sync.WaitGroup
	sem := make(chan int, maxConcurrentWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		sem <- 1 // Chiếm một slot
		go worker(i, &wg, sem)
	}

	wg.Wait()
	fmt.Println("All workers are done")
}
