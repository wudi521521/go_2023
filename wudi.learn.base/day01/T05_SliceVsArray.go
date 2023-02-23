package main

import "fmt"

type P struct {
	name string
}

//go语言中国slice map interface channel 默认使用引用传递
func m0(v [3]int){
	v[0] = 8
}

func m1(v []int){
	v[0]=8
}

func m2(p P){
	//这个地方的定义：是复制一份，在复制的target上进行了修改lisi，原有的source没有修改
	p.name = "lisi"
}

func m3(p *P){
	//传递的是地址，修改地址中的属性
	p.name = "lisi"
}
func main() {
	var v1 = [] int{1,2,3}
	fmt.Println(v1)

	//Object creat and print
	var pp P = P{"zhang san"}
	fmt.Println(pp)
	m2(pp)
	fmt.Println(pp)
	m3(&pp)
	fmt.Println(pp)



}
