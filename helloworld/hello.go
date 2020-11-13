package main

import (
	"fmt"

	"github.com/captainlee1024/stringutil"
)

func main() {
	fmt.Printf("Hello World!")
	fmt.Printf(stringutil.Reverse("Hello World!你好！") + "\n")
	fmt.Println("Hi gopher!")
}
