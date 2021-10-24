package main

import "fmt"

func main() {
    //var price int32 = -68
    //fmt.Println(price)
    //var price2 uint32 = 68
    //fmt.Println(price2)
    //
    //var price3 float64 = 99.99
    //fmt.Println(price3)
    //
    //var checked bool = false
    //fmt.Println(checked)
    //
    //var name string = "从0到Go语言微服务架构师-学习路线图"
    //fmt.Println(name)
    //
    ////byte
    //fmt.Println([]byte(name))

    //rune对标char

    //complex()

    name,price ,num := "红烧肉",68,1
    var total = 0
    discount:= 0.7
    total =  int(float64(price) * float64(num) * discount)
    fmt.Printf("%s总价%d\n",name,total)
}
