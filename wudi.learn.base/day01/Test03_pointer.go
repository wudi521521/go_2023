package main

import "fmt"

func ff(v *int){
	*v = 8//解引用
}

func main() {
    var i int = 8
    fmt.Println(&i)

    var pi *int
    pi = &i
    fmt.Println(pi)

    m :=0
    ff(&m)
    fmt.Println(m)

    n :=0
    fmt.Println(&n)

}
