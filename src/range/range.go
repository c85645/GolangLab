package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for i, value := range pow {
		fmt.Printf("2^%d = %10d = %s\n", i, value, fmt.Sprintf("%010b", value))
	}
}
