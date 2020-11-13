package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "yes我爱中国!"
	fmt.Println(s, len(s))
	for i, b := range []byte(s) {
		fmt.Printf("%d %X\n", i, b)
	}
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}

	fmt.Println("Rune count",
	utf8.RuneCountInString(s))	


	bytes := []byte(s) // 拿到 s 的字节
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes); // 使用 DecodeRune 拿到一个一个的字符，返回值是 rune 和 字符的 zise
		bytes = bytes[size:] //从下一个字符开始解码
		fmt.Printf("%c ", ch) // 把字符打印出来
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}

}