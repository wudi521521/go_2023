package main

import "fmt"

func main() {
	var v *int
	fmt.Println(*v)
	fmt.Println(v)
	v = new(int)
	*v = 8
	fmt.Println(*v)

}
