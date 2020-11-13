package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
		case "+":
			return  a + b, nil
		case "-":
			return  a - b, nil
		case "*":
			return  a * b, nil
		case "/":
			q, _ := div(a, b)
			return  q, nil
		default:
			// panic("unsupported operation : " + op)
			return 0, fmt.Errorf(
				"unsupported operation : %s", op)

	}
}

// func div(a, b int) (int, int) {
// 	return a/b, a%b
func div(a, b int) (q, r int) {
	q = a/b
	r = a%b
	return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
    opName := runtime.FuncForPC(p).Name()
    fmt.Printf("Calling function %s with args " +
               "( %d %d )", opName, a, b)
    return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func printArray(arr [3]int){
	// 遍历
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 使用 range 进行遍历
	for i := range arr{
		fmt.Println(arr[i])
	}

	// 获取值
	for _, v := range arr {
		fmt.Println(v)
	}

	// 获取值和下标
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	fmt.Println(
		eval(3, 4, "x"),
	)
	fmt.Println(div(3, 4))
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a),float64(b)))
		}, 3, 4,
	))
	fmt.Println(sum(1, 2, 3))

	var arr1 [3]int
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{4, 5, 6}
	var grid [4][5]bool
	fmt.Println(arr1, arr2, arr3, grid)
	printArray(arr1)
	printArray(arr3)
}