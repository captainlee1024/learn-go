package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g() // 获取下一个元素
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small.

	return strings.NewReader(s).Read(p)
}

func printFileContents(read io.Reader) {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

type fib func() (int, int, fib)

func fibonacci2(a, b int) fib {
	return func() (int, int, fib) {
		a, b = b, a+b
		return a, b, fibonacci2(a, b)
	}
}

func main() {
	f := fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
	printFileContents(f)

	f2 := fibonacci2(0, 1)
	for i := 0; i < 10; i++ {
		a, _, _ := f2()
		fmt.Println(a)
	}
}
