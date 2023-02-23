package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//定义一个整数类型
	var num1 int8 = 120
	fmt.Println(num1)

	var num2 uint8 = 200
	fmt.Println(num2)

	var num3 = 28
	//Prinf函数的作用就是：格式化的，
	fmt.Printf("%T",num3)
	fmt.Println()
	//字节数
	fmt.Println(unsafe.Sizeof(num3))
}
