package main

import "fmt"

func main() {
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 10)
	fmt.Println(result)
}
