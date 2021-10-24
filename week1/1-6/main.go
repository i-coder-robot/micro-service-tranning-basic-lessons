package main

import (
	"fmt"
	"net/http"
)

func GoToHappy(sunShine bool){
	if sunShine {
		fmt.Println("出去嗨皮,出去浪")
	}else{
		fmt.Println("从0到Go语言微服务架构师-路线图")
	}
}

func VisitUrl(url string) (int,error){
	res,err:=http.Get(url)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return 0,err
	}else{
		fmt.Printf("%d",res.StatusCode)
		return res.StatusCode,nil
	}
}

func SwitchShow(url string)  {
	if code,err:=VisitUrl(url);err!=nil{
		fmt.Println(err)
	}else{
		switch code {
		case 200:
			fmt.Println("\n请求成功")
		case 404:
			fmt.Println("\n网址不存在")
		default:
			panic("未知错误")
		}
	}
}

func SwitchShow2(url string)  {
	if code,err:=VisitUrl(url);err!=nil{
		fmt.Println(err)
	}else{
		switch {
		case code==200:
			fmt.Println("\n请求成功")
		case code == 404:
			fmt.Println("\n网址不存在")
		default:
			panic("未知错误")
		}
	}
}

func main() {
	//VisitUrl("http://www.baidu.com")
	SwitchShow("http://www.baidu.com")
}
