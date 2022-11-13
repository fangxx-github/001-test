package main

import (
	"fmt"
	"reflect"
)

func main() {
	var money interface{} = 10
	//var money interface{} = 10.0
	//var money interface{} = "10"
	switch money.(type){
	case int:
		fmt.Println("money is int")		
	case int64:
		fmt.Println("money is int64")
	case float64:
		fmt.Println("money is flost64")
	default:
		fmt.Println("monet is weizhi")
	}
	fmt.Println("END",money, money,reflect.TypeOf(money))
	money = "10"
	fmt.Println("END",money, money,reflect.TypeOf(money))
}