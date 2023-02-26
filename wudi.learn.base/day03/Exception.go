package main

import "fmt"

func main() {
	//引入机制：defer+recover机制处理错误
	test03()
	fmt.Println("成功")
}

func test03(){
	defer func() {
		//调用recover内置函数，可以捕获错误
		err := recover()
		//如果没有捕获错误，返回值为零值：nil
		if err != nil {
			fmt.Println("错误已经捕获")
			fmt.Println("err",err)
		}
	}()
	num1 :=10
	num2 :=0
	result := num1/num2
	fmt.Println(result)
}