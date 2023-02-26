package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 int = 20
	exchangeNum(num1, num2)
	fmt.Println(num1,num2)

}

func exchangeNum(num1 int,num2 int)  {
	var t int
	t = num1
	num1 = num2
	num2 = t
	fmt.Println(num1,num2)
}
