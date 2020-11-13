package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/captainlee1024/learngo/functional/fib"
)

func tryDefer() {
	// defer fmt.Println(1)
	// fmt.Println(2)
	// defer fmt.Println(3)
	// // panic("err")
	// return
	// // defer fmt.Println(4)
	// // fmt.Println(5)

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i > 10 {
			panic("printed to many")
		}
	}
}

func writerFile(filename string) {
	// file, err := os.Create(filename)
	file, err := os.OpenFile(
		filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a customer error")
	if err != nil {
		// panic(err)
		// fmt.Println("file already exists")
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	// 直接写文件内比较慢，使用 bufio 先写到内存里，在一次性写入文件速度比较快
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()

	writerFile("fib.txt")

}
