package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) *
				time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	// var c1, c2 chan int
	var c1, c2 = generator(), generator()
	// w := createWorker(0)
	var worker = createWorker(0)

	// n := 0
	// hasValue := false
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			// hasValue = true
			// fmt.Println("Received from c1:", n)
			// w <- n
			values = append(values, n)
		case n := <-c2:
			// hasValue = true
			// fmt.Println("Received from c2:", n)
			// w <- n
			values = append(values, n)
		// default:
		// 	fmt.Println("No value received")
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len = ", len(values))
		case <-tm:
			fmt.Println("bye bye")
			return
		}
	}
}
