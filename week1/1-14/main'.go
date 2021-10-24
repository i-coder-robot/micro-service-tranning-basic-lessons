package main

import (
	"fmt"
	"mic-basic-lessons/week1/1-14/hello"
)

func main() {
	//var Food model.Food
	//fmt.Println(Food)

	price:=68
	newPrice := model.MyInt(price)
	fmt.Println(newPrice)

	food := model.Food{Name: "豆汁儿"}
	p:= model.Person{Food: food}
	p.SayHello("老张")

	food2 := model.Food2{
		FoodName:   "葱烧海参",
		FoodSystem: model.FoodSystem{Name: "鲁菜"},
		Another:    model.FoodSystem{Name: "鲁菜"},
	}
	fmt.Println(food2.Another.Name)
	fmt.Println(food2.Name)
}
