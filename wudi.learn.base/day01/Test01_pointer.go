package main

import "fmt"

func f(v* int){
	*v = 8
}
func main() {
 var i int =8
 fmt.Println(&i)

 var pi*int
 pi = &i
 fmt.Println(pi)

 m :=0
 f(&m)
 fmt.Println(m)

 n :=0
 fmt.Println(&n)
}
