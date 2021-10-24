package main

import (
	"fmt"
	"time"
)

func main() {
	//sum :=0
	//for i:=0;i<100;i++{
	//	sum+=i
	//}

	i:=0
	//for ;i<100;i++{
	//	sum+=i
	//}

	//for ;;i++{
	//	if i==100{
	//		break
	//	}
	//	sum+=i
	//}
	//for ;;{
	//	if i==100{
	//		break
	//	}
	//	i++
	//	sum+=i
	//}
	//fmt.Println(sum)
	for{
		fmt.Println("从0到Go语言微服务架构师",i)
		i++
		time.Sleep(time.Duration(3)*time.Second)
	}

}
