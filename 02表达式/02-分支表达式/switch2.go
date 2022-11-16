package main

import (
	"fmt"

)

func main() {
	var money int
	var busy bool
	switch {
	case money >= 0 && money <= 20:
		fmt.Println("点个外卖")
		if busy {
			break
		}
		fmt.Println("在走走把")
		//fallthrough
	case money >= 20 && money <= 2000:
		fmt.Println("下馆子")
	default:
		fmt.Println("在看看")
	}
	fmt.Println("END")
}
//分支表达式
//1.可随时终止分支表达式
//2.并入其他条件分支继续执行
//3.原生支持类型判断
//