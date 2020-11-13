package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	b := [5]int{1: 10, 3: 20, 0: 30}

	fmt.Println(a)
	fmt.Println(b)
}
