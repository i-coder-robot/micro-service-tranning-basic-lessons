package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday

	 Author = "欢喜"
	 Book = "Go语言极简一本通"
)

func main() {
	const version="1.0"
	const appName = "面向加薪学习"

	fmt.Printf("%s-%s-%s-%s\n",Author,Book,version,appName)

	fmt.Println(Friday)
}
