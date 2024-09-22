package main

import "fmt"

// 簡單的一個函式，實現了參數+1 的操作
func add1(a *int) int {
	*a = *a + 1 // 我們改變了 a 的值
	return *a   // 回傳一個新值
}

func main() {
	x := 3

	fmt.Println("x = ", x) // 應該輸出 "x = 3"

	x1 := add1(&x) //呼叫 add1(x)

	fmt.Println("x+1 = ", x1) // 應該輸出"x+1 = 4"
	fmt.Println("x = ", x)    // 應該輸出"x = 3"
}
