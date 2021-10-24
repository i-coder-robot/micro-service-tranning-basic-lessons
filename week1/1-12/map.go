package main

import "fmt"

func main() {
	m := make(map[string]string)
	fmt.Println(m)

	var m2 map[string]string
	fmt.Println(m2)

	m3 := map[string]string{
		"鲁菜": "九转大肠",
		"川菜": "麻婆豆腐",
	}
	fmt.Println(m3)

	m4 := map[string]map[int]string{}
	fmt.Println(m4)

	m5 := map[int]string{}
	m5[1] = "葱烧海参"
	m5[2] = "回锅肉"
	for idx, item := range m5 {
		fmt.Println(idx, item)
	}

	for _, item := range m5 {
		fmt.Println(item)
	}

	//key ：chiken apple orane

	v:=m5[1]
	fmt.Println(v)

	v2,ok:=m5[5]
	if ok {
		fmt.Println(v2)
	}else{
		fmt.Println("找不到元素")
	}

	//delete
	//
	delete(m5,1)
	//fmt.Println(m5)

	delete(m5,10)
	fmt.Println(m5)
}
