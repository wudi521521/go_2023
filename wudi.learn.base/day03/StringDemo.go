package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//len
	 str := "golang 你好"//在golang中，汉字是utf-8字符集，一个汉字3个字节
	fmt.Println(len(str))

	 //切片是9
	 r := []rune(str)
	 fmt.Println(len(r))
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n",r[i])

	}
	fmt.Println(strconv.Atoi("111"))
	fmt.Println(strconv.Itoa(99))
	fmt.Println(strings.Count(str,"你好"))
	fmt.Println(strings.EqualFold("d","D"))

}
