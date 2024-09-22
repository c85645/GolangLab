package main

import "fmt"

// 宣告一個新的型別
type person struct {
	name string
	age  int
}

// 比較兩個人的年齡，回傳年齡大的那個人，並且回傳年齡差
// struct 也是傳值的
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age { // 比較 p1 和 p2 這兩個人的年齡
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person

	// 賦值初始化
	tom.name, tom.age = "Tom", 18

	// 兩個欄位都寫清楚的初始化
	bob := person{age: 25, name: "Bob"}

	// 按照 struct 定義順序初始化值
	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)
}
