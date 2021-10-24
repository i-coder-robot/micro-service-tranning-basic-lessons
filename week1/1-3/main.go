package main

import "fmt"

var (
  book = "\nGo语言极简一本通"
  lesson1 = "\nGo语言微服务架构核心22讲"
  lesson2 = "\n从0到Go语言微服务架构师路线图"
)

func main() {
  // int price = 68
  // var 变量名称 变量类型
  var price int32 = 68
  var name = "宫保鸡丁"
  fmt.Println("hello world")
  fmt.Printf("%d,%s",price,name)

  var price2,weight = 68,1
  fmt.Printf("\n%d,%d",price2,weight)

  price3:=68
  weight2 :=1
  fmt.Printf("\n%d,%d",price3,weight2)

  fmt.Println(book)
  fmt.Println(lesson1)
  fmt.Println(lesson2)
}
