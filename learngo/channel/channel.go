package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		// n, ok := <-c
		// if !ok {
		// 	break
		// }
		// fmt.Printf("worker %d received %c\n", id, <-c)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	// go func() {
	// 	for {
	// 		fmt.Printf("worker %d received %c\n", id, <-c)
	// 	}
	// }()
	go worker(id, c)
	return c
}

func chanDemo() {
	// var c chan int // c == nil, nil 的 channel  没有办法使用
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		// channels[i] = make(chan int)
		// go worker(i, channels[i])
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Second)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond * 2)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	// close(c)
	time.Sleep(time.Millisecond * 3)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
