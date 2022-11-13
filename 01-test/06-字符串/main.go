//格式化输出   http://doc.golang.ltd/

package main

import "fmt"

func main() {
	fmt.Println("中文，数字，符号，❤")
	fmt.Printf("我的名字：%s\n", "xiao")
	fmt.Printf("我的名字：%q\n", "xiao")
	fmt.Printf("我的名字：%x\n", "xiao")
	fmt.Printf("我的名字：%X\n", "xiao")
}
