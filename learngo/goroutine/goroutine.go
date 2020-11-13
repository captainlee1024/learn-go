package main

import (
	"fmt"
	"time"
)

func main() {
	// for i := 0; i < 1000; i++ {
	// 	go func(i int) {
	// 		for {
	// 			fmt.Printf("hello from "+
	// 				"groutine %d\n", i)
	// 		}
	// 	}(i)
	// }
	// time.Sleep(time.Millisecond)

	var a [10]int
	// s := 'a'
	for i := 0; i < 10; i++ {
		// s = s + 1
		go func(i int) {
			for {
				a[i]++
				// runtime.Gosched()
				// fmt.Println(s)
				fmt.Println(i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond*3)
	// fmt.Println(a)
}
