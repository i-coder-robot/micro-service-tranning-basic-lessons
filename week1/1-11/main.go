package main

import "fmt"

func SoldOut(foods []string){
	foods[1] = "已售罄"
	fmt.Println(foods)
}

func main() {
	array5:= [...]string{"烧乳猪","鲍鱼","文昌鸡","大龙虾","烧鹅","茶田鸭"}
	//Slice View 视图
	slice1 := array5[1:3]
	fmt.Println(slice1)

	slice2 := array5[1:]
	fmt.Println(slice2)

	slice3 :=array5[:3]
	fmt.Println(slice3)

	slice4:=array5[:]
	fmt.Println(slice4)

	//SoldOut(slice4)
	//fmt.Println(slice4)

	slice5 := slice4[2:5]
	fmt.Println(slice5)

	slice6:= slice5[2:4]
	fmt.Println(slice6)

	//slice7:= slice5[2:8]
	//fmt.Println(slice7)
	slice8:=[]string{}
	fmt.Println(len(slice8),cap(slice8))
	for i:=0;i<10;i++{
		slice8 = append(slice8,fmt.Sprintf("蛋炒饭%d",i))
		fmt.Println(len(slice8),cap(slice8))
	}

	//扩容--聚会，人多，换大桌。

	s1:=make([]string,8)
	s1=append(s1, "欢喜")
	s1=append(s1, "666")

	s2:=make([]string,8,16)
	s2=append(s2,"GO语言极简一本通")




	//copy
	copy(s1,slice4)
	copy(s2,slice4)
	fmt.Println(s1)
	fmt.Println(s2)

	//delete
	//slice4[1:]
	fmt.Println(slice4[:len(slice4)-1])
}
