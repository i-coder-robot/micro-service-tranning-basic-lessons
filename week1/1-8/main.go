package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

//func 函数名称(参数列表) 返回值列表

//public String 名称(参数列表)

func ShowName() string{
	return "从0到Go语言微服务架构师"
}

func ShowInfo(bookName,author string) string {
	return bookName + "-作者是" + author
}

func ShowInfoAndPrice(bookName,author string,price float64) (string,float64) {
	return bookName + "-作者是" + author,price
}

func ShowInfoAndPrice2(bookName,author string,price float64) (bookInfo string,finalPrice float64) {
	bookInfo = bookName + "-作者是" + author
	finalPrice= price
	return
}

func VisitUrl(url string) (int,error){
	res,err:=http.Get(url)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return 0,err
	}else{
		return res.StatusCode,nil
	}
}

func SwitchShow2() (int,error){
	if code,err:=VisitUrl("http://www.baidu.com");err!=nil{
		return 0,err
	}else{
		return code,nil
	}
}

func PrintBookInfo(do func(string,string,float64)(bookInfo string, finalPrice float64),
	bookName,author string,price float64){
	pointer:=reflect.ValueOf(do).Pointer()
	methodName:=runtime.FuncForPC(pointer)
	fmt.Println(methodName)
	bookInfo,finalPrice := do(bookName,author,price)
	fmt.Println(bookInfo)
	fmt.Println(finalPrice)
}

func PrintBookInfo2(bookName,author string,price float64){
	do := func(string,string,float64)(bookInfo string, finalPrice float64){
		bookInfo = bookName + "的作者是-" + author
		finalPrice= price
		return
	}
	fmt.Println(do(bookName,author,price))
}

func ShowAll(foodList ...string) string{
	r:=""
	for idx,item:=range foodList {
		fmt.Println(idx)
		r+=item
	}
	return r
}

func main() {
	//fmt.Println(ShowName())
	//info,_:= ShowInfoAndPrice2("Go语言极简一本通","欢喜",99.00)
	//fmt.Println(info)

	//fmt.Println(SwitchShow2())
	//PrintBookInfo(ShowInfoAndPrice2,"Go语言极简一本通","欢喜",99.00)
	//PrintBookInfo2("Go语言极简一本通","欢喜",99.00)

	fmt.Println(ShowAll("葱烧海参","九转大肠","糖心鲍鱼","麻辣兔头"))
}
