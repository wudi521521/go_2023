package main

import "fmt"

func main() {
fmt.Println(demoName(1,2))
}

func demoName (num1 int,num2 int)(result int,num3 int){//如果返回值类型就是一个的话，那么（）是可以省略不写
	return num1+num2,0
}
