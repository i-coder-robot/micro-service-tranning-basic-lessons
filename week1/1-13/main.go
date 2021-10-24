package main

import "fmt"

//interface 契约

type Food struct {
	No int32
	Name string
}

func NewFood(no int32,name string) *Food{
	return &Food{
		No: no,
		Name: name,
	}
}

func (f *Food) Show(){
	fmt.Println(f.Name)
}

func (f *Food) SetName(name string){
	f.Name = name
}

//func Show(f *Food){
//	fmt.Println(f.Name)
//}

func main() {
	var food Food
	food = Food{
		No: 1,
		Name: "葱烧海参",
	}
	fmt.Println(food)
	fmt.Println(food.No)
	fmt.Println(food.Name)

	food2:=Food{
		No:   2,
		Name: "麻婆豆腐",
	}
	fmt.Println(food2)

	//Food food= new Food()

	food3:=NewFood(3,"回锅肉")
	fmt.Println(food3)

	food3.SetName("宫保鸡丁")
	fmt.Println(food3)
}
