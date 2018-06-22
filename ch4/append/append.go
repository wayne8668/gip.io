package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y

	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
	fmt.Println("---------------------")
	a := make([]int,3,3)
	b := []int{1,2,3,4,5,6,7,8,9}
	//copy(dest,src)
	//copy函数不会改变dest slice的长度和容量，只是slice元素对应下标值的拷贝，非底层数组的引用。函数返回拷贝的元素个数(实际为两个slice长度的最小值)
	c := copy(a,b)
	fmt.Println("a's len is:", len(a))
	fmt.Println("a's cap is:", cap(a))
	//因为copy不是底层数据的引用，只是元素赋值，因此改变dest的元素值，不会影响src的值
	a[0] = -1
	fmt.Println(b[0])	//这里值保持不变，还是0
	fmt.Println(a,b,c)
}
