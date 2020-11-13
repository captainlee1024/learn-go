package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	// Creating slice
	var s []int //Zero value for slice nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	// fmt.Println(s)
	// s1 := []int{1, 2, 3, 4}
	// printSlice(s1)

	// s2 := make([]int, 16)
	// printSlice(s2)

	// s3 := make([]int, 10, 32)
	// printSlice(s3)

	// // Copying slice
	// copy(s2, s1)
	// printSlice(s2)
	// printSlice(s1)

	// // Deleting elements form slice
	// s2 = append(s2[:3], s2[4:]...)
	// printSlice(s2)
	// // Popping from font
	// front := s2[0]
	// s2 = s2[1:]
	// fmt.Println(front)
	// printSlice(s2)
	// //Popping from back
	// tail := s2[len(s2)-1]
	// s2 = s2[:len(s2)-1]
	// fmt.Println(tail)
	// printSlice(s2)
}
