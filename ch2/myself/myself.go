package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Println(s) //输出string默认"""

	var i, j, k int                  //默认赋0值
	var b, f, g = true, 2.3, "hello" //给3个类型的变量赋值

	fmt.Printf("%t %f %s\n", b, f, g)
	fmt.Printf("%d %d %d\n", i, j, k)

	x := 1
	p := &x //&x为x变量的指针
	fmt.Println(p)
	fmt.Println(*p)
	*p = 2 //*p 等价于 x，*访问p指针的值
	fmt.Println(x)

	m, n := 0, 1
	fmt.Printf("%d %d\n", m, n)
	m, n = n, m
	fmt.Printf("%d %d\n", m, n)

	fmt.Println("-----------------------------------------------------")

	d := new(int)   //new函数返回对应类型的指针，int类型返回默认0值
	fmt.Println(d)  //输出地址
	fmt.Println(*d) //输出值变量，默认值0
	*d = 2          //赋值2
	fmt.Println(*d) //输出值变量2

	fmt.Println(delta(3, 2))

	fmt.Println("-----------------------------------------------------")

	v := 1
	v++ //等价于v = v + 1 ; 然而GO 并没有++v和--v

	e,h := 1,2
	fmt.Printf("e = %d ; h = %d\n",e,h)
	//多重赋值，允许几个变量一次性被赋值
	e,h = h,e	//交换变量值，编译器会一次完成赋值。在变量同时出现在赋值符两侧的时候特别有用
	fmt.Printf("e = %d ; h = %d\n",e,h)
	//交换变量值是一步完成，并不等价于下面代码。下面分两步，然后那样并不能完成交换变量值
	e = h
	h = e
	fmt.Printf("e = %d ; h = %d\n",e,h)

	r,s := 1,"ad"
	fmt.Printf("r = %d ; s = %s\n",r,s)
}

func delta(old, new int) int { //new 在GO里是一个内置函数，不是一个关键字，因此这里参数列表里是可以用new作为参数名传入的
	// i := new(int)			  //一旦函数列表中使用了new作为参数名，new 函数就不可以使用了。
	// return new - old + *i

	//  i := new(int)		 //如果把函数列表中的参数new换成newa作为变量名，new 函数就又可以使用了。
	//  return newa - old + *i

	return new - old
}
