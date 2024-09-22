package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1 // 初始化斐波那契數列的前兩個數字
	for {
		select {
		case c <- x: // 將 x 發送到 channel c
			// 生成下一個斐波那契數
			x, y = y, x+y
		case <-quit: // 如果從 quit channel 接收到數據，則退出
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)    // 創建一個整數 channel c
	quit := make(chan int) // 創建一個整數 channel quit
	go func() {            // 啟動一個匿名 Goroutine
		for i := 0; i < 10; i++ { // 打印 10 個斐波那契數字
			fmt.Println(<-c) // 從 channel c 接收數字並打印
		}
		quit <- 0 // 完成後向 quit channel 發送信號，結束 fibonacci 函數
	}()
	fibonacci(c, quit) // 主 Goroutine 執行 fibonacci 函數
}
