package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // 隨機延遲
	c <- total                                                   // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
