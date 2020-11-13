package main

import (
	"fmt"
	"sync"
)

// type atomicInt int
type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	// *a++
	a.lock.Lock()
	defer a.lock.Unlock()

	a.value++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func f1() int {
	x := 5
	x++
	defer func() {
		fmt.Println(x)
		x++
		fmt.Println(x)
	}()
	return x
}

func f2() (x int) {
	x = 5
	defer func() {
		x++
	}()
	return
}

func a() (i int) {

	fmt.Println(i)
	i++
	fmt.Println(i)
	defer fmt.Println(i)
	return 1
}

func main() {
	// var a atomicInt
	// a.increment()
	// go func() {
	// 	a.increment()
	// }()
	// time.Sleep(time.Millisecond)
	// fmt.Println(a.get())

	// 闭包 defer和return执行过程
	fmt.Println(f1())
	// x := f2()
	// fmt.Println(x)
	// fmt.Println(a())
}
