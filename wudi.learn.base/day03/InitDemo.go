package main

import "fmt"

var num int = test01()

func test01() (int) {
	fmt.Println("test func running")
	return 10
}

func init(){
	fmt.Println("init func running")
}

func main() {
fmt.Println("main running")
}
