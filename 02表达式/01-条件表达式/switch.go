package main

import (
	"fmt"

)

func main() {
	var money int
	switch money {
	case 20:
		fmt.Println("点个外卖")
	case 200:
		fmt.Println("下馆子")
	case 2000:
		fmt.Println("米其林")
	default:
		fmt.Println("想象")
	}
	fmt.Println("end")
}



//可随时中断
//可扭转到下一个 fallthrough 