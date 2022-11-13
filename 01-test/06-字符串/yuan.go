package main

import "fmt"

func main() {
	var a float64 = 1.5
	var b float64 = 2.5
	const pi float64 = 3.1415926
	mianji := pi * (b * b) + pi * (a * a)
	fmt.Println("a和b⚪的面积和: %.3f", mianji)
}
