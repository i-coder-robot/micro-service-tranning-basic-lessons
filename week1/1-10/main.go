package main

import "fmt"
//[]string
func printFood(foods [3]string){
 foods[2]="大饼"
 fmt.Println(foods)
}

func printFood2(foods *[3]string){
 foods[2]="大饼"
 fmt.Println(foods)
}

func main() {
 var array1 [6]string
 array2 :=[3]string{"火锅","烧烤","家常菜"}
 array3 := [...]string{"火锅","烧烤","家常菜"}
 fmt.Println(array1)
 fmt.Println(array2)
 fmt.Println(array3)
 var matrix [3][4]string
 var matrix2 [3][4]int
 fmt.Println(matrix)
 fmt.Println(matrix2)
 for i:=0;i<len(array2);i++{
  fmt.Println(array2[i])
 }

 for idx,item :=range array2{
  fmt.Println(idx,item)
 }

 for _,item :=range array2{
  fmt.Println(item)
 }
 //java
 //for (类型 变量 :集合){}
 //python
 //for 变量 in 集合:
 // 语句块

 printFood2(&array2)
 fmt.Println(array2)
}
