package main

import (
	"fmt"
	"sync"
)

func doworker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		// go func() {
		// 	done <- true
		// }()
		// done <- true
		// wg.Done()
		w.done()
	}
}

type worker struct {
	in chan int
	// done chan bool
	// wg *sync.WaitGroup
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doworker(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20) // 当我们知道有多少个任务时
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		// <-workers[i].done
	}
	// for _, worker := range workers {
	// 	<-worker.done
	// 	// <-worker.done
	// }
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		// <-workers[i].done
	}

	//wait for all of them
	// for _, worker := range workers {
	// 	<-worker.done
	// 	// <-worker.done
	// }

	wg.Wait()
}
func main() {
	chanDemo()
}
