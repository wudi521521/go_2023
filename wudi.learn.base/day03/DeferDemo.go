package main

import "fmt"

func main() {
fmt.Println(add(30,50))
}

func add(num1 int, num2 int)(int){
	//在golang中，程序遇见defer关键字，不会立即执行defer后的语句，而是将defer后的语句押入一个栈中，然后继续执行函数后面的。
	defer fmt.Println("num1 = ",num1)
	defer fmt.Println("num2 = ",num2)
   num1 +=90
   num2 +=100
	//栈的的特点是先进后出
	//在函数执行完毕以后，从栈中取取出语句开始执行，按照先进后出的规则执行语句
	var sum int = num1 +num2
	fmt.Println("result",sum)
	return sum
}
