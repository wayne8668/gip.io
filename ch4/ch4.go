package main

import (
	"fmt"
)

const blen = 10

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	a := [...]int{11, 12, 13, 14, 15, 16, 17, 18, 19} //数组长度由赋值的数量决定
	b := [blen]int{1, 2, 3, 4, 5, 6, 7, 8, 9}         //赋值不够补对应类型的0值

	fmt.Println("a's len is:", len(a)) //len函数返回数组长度
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("a's first item:", a[0])
	fmt.Println("a's last item:", a[len(a)-1])

	//遍历数组
	for k, v := range a {
		fmt.Printf("index:%d value:%d\n", k, v)
	}

	//仅输出元素
	for _, v := range a {
		fmt.Printf("value:%d\n", v)
	}
	fmt.Println("--------------------")

	//USD,EUR,GBP,RMB 是常量，取值为0，1，2，3，对应数组的下标
	symbol := [...]string{USD: "$", EUR: "E", GBP: "F", RMB: "Y"}
	fmt.Println(symbol[RMB])

	for k, v := range symbol {
		fmt.Printf("index:%d value:%s\n", k, v)
	}

}
