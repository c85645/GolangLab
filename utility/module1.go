package utility

import "fmt"

type testInt func(int) bool // 宣告了一個函式型別

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 宣告的函式型別在這個地方當做了一個參數

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // 函式當做值來傳遞了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven) // 函式當做值來傳遞了
	fmt.Println("Even elements of slice are: ", even)
}
