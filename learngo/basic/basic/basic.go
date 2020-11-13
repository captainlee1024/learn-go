package main

import (
	"fmt"
	"math"
)

var a int
var c = 2
var s = "def"
var b = true

var (
	aa = 1
	cc = 3
	ss string
	bb bool
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d %s\n", a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 1, 2
	var s string = "s初始值"

	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 1, 2, true, "String-字符串"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 1, 2, true, "初始值-String"
	b = 3
	fmt.Println(a, b, c, s)
}

func triangle() {
	a, b := 3, 4
	// var c int
	// c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const a, b = 3, 4
	const filename = "abc.txt"
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		// cpp = 0
		// java = 1
		// python = 2
		// golang = 3
		cpp = iota
		java
		python
		golang
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	fmt.Println(a, c, s, b, aa, cc, ss, bb)
	var a int
	var b bool
	var s string
	fmt.Println(a, b, s)
	fmt.Printf("%d %t %q", a, b, s)
	triangle()
	consts()
	enums()
}
