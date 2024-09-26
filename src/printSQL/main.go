package main

import (
	"fmt"
	"os"
)

func main() {
	start := 94107
	end := 94610

	// 打開或創建DROP.sql檔案
	file, err := os.Create("DROP.sql")
	if err != nil {
		fmt.Println("無法創建檔案:", err)
		return
	}

	for i := start; i <= end; i++ {
		// 將數字轉為5位數的格式，例如：00001
		tableSuffix := fmt.Sprintf("%05d", i)
		// 生成 DROP TABLE 語句
		dropStatement := "DROP TABLE SNS_CAMP_" + tableSuffix + ";"
		// 寫入檔案
		_, err := file.WriteString(dropStatement)
		if err != nil {
			fmt.Println("寫入檔案時出錯:", err)
			return
		}
		fmt.Println(dropStatement)
	}
	fmt.Println("DROP.sql 檔案已成功生成！")
}
