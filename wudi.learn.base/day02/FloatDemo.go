package main

import "fmt"

/*
*float32 4个字节
 float64 8个字节
 ps：底层存储空间和操作系统无关
 ps：浮点类型底层存储：符合位+指数位+尾数位，所以尾数位只是存了一个大概，很可能出现精度的损失
 */
func main() {
//定义浮点类型
	var num01 float32 = 3.14;
	fmt.Println(num01)
	//存储单个字符，一般使用byte
	//定义字符类型的数据
	var c1 byte = 'a'
	fmt.Println(c1)
	var c2 byte = '6'
	fmt.Println(c2)

	var c3 byte = '('
	fmt.Println(c3)
	//字符类型，本质上就是一个整数，也可以直接参与运算
	fmt.Printf("%c",c3)
}
