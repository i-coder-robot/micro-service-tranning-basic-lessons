package model

import "fmt"

func Cook(f Food) {}

func Send(f Food) {}

type MyInt int32

type Hello func(string)

type Person struct {
	Food Food
}

type Food struct {
	Name string
}
func (p *Person) SayHello(name string) {
	fmt.Println(name+"吃了嘛？"+p.Food.Name)
}

type FoodSystem struct {
	Name string
}
type Food2 struct {
	FoodName string
	FoodSystem
	Another FoodSystem
}