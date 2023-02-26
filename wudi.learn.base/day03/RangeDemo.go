package main

import "fmt"

/**
for range 结构是Go语言特有的一种迭代结构，在许多情况下都非常有用，for range可以遍历数组，切片
，字符串，map及其通道。
for range语法上类似于其他
for key ，val = range coll
 */
func main() {
   //定义一个字符串
	var str string = "hello golang"
	//普通for循环,按照字节进行循环
	for i :=0;i<len(str);i++{
		fmt.Printf("%c \n",str[i])
	}
	//for range
	for i, value := range str {
		fmt.Printf("索引为：%d，具体值:%c \n",i,value)
		if i==3 {
			fmt.Println("终止")
			//停止当前循环，break的作用结束离它最近的循环
			break
		}

	}
	return
	fmt.Println("return")
	//fmt.Println(i,value)
}
