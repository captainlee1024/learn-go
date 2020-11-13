package queue

import "fmt"

func ExampleQueue_Popo() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Popo())
	fmt.Println(q.Popo())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Popo())
	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 2
	// false
	// 3
	// true
}
