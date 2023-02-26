package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间now
	now := time.Now()
	//Now函数返回值是一个结构体：2023-02-26 17:29:25.727961 +0800 CST m=+0.002507282
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Date())

	fmt.Println("-------------")
	//必须这样写
	datestr2 := now.Format("2006/01/02 15/04/05")
	fmt.Println(datestr2)


}
