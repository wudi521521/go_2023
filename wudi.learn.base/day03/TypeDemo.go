package main

import "fmt"

func test(num int){
fmt.Println(num)
}

func main() {
    a := test
    fmt.Printf("%T,%T",a,test)//func(int),func(int)
    //通过该变量可以对函数调用
    a(10)//等价于test(10)

    //自定义数据类型（相当于起别名）
	type myInt int
    var num1 myInt = 30
    fmt.Println(num1)
    var num2 int = 30
    //虽然是别名，但是在go中编译识别的时候还是认为myInt和int不是一种类型
    num2 = int(num1)
    fmt.Println("num2",num2)

}
