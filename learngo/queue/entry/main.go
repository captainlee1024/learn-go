package main

import (
	"fmt"

	"github.com/captainlee1024/learngo/queue"
)

func main() {
	q := queue.Queue{1}

	q.PushInt(2)
	q.PushInt(3)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.PopoInt())
	fmt.Println(q.PopoInt())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.PopoInt())
	fmt.Println(q.IsEmpty())

	q.Push("字符串abc")
	fmt.Println(q.Popo())
}
