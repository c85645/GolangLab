package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[1:4]
	printSlice(s) // 3, 5, 7

	s = s[:2]
	printSlice(s) // 3, 5

	s = s[1:]
	printSlice(s) // 5
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
