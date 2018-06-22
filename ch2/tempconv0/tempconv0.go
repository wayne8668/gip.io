package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//toString 方法
func (c Celsius) String() string { return fmt.Sprintf("%g *C", c) }

/*
* 同命名类型基于底层类型进行运算/比较
 */
func main() {
	//同命名类型基于底层类型进行运算/比较
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	fmt.Println(boilingF == CToF(FreezingC))
	// fmt.Printf("%g\n", boilingF-FreezingC)	//编译错误，类型不匹配——不同类型无法进行运算和比较

	var c Celsius
	var f Fahrenheit
	//命名类型基于底层类型直接与底层类型进行运算/比较
	fmt.Println("'c == 0' ->", c == 0)
	fmt.Println("'f >= 0' ->", f >= 0)
	// fmt.Println(c == f)	//编译错误，类型不匹配——不同类型无法进行运算和比较
	fmt.Println("'c == Celsius(f)'->", c == Celsius(f))
	fmt.Println(Celsius(1000))
	fmt.Println("---------------------")

}
