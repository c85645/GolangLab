package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") //開一個新的 Goroutines 執行
	say("hello")    //當前 Goroutines 執行
}

// 以上程式執行後將輸出：
// hello
// world
// hello
// world
// hello
// world
// hello
// world
// hello
