package main

import (
	"fmt"
)

func main() {
	//声明初始化定长为13的string数组，months[0]为string初始值
	months := [...]string{1: "yiyue", 2: "eryue", 3: "sanyue", 4: "siyue", 5: "wuyue", 6: "liuyue", 7: "qiyue", 8: "bayue", 9: "jiuyue", 10: "shiyue", 11: "shiyi", 12: "shier"}
	//months[0]为空字符串
	fmt.Println(months[0])
	//声明初始化slice Q2
	Q2 := months[4:7]
	//Q2长度为item个数7-4=3
	fmt.Println("Q2's len is:", len(Q2))
	//Q2容量为len(months)减去months从下标4开始截取的位置为13-4=9
	fmt.Println("Q2's cap is:", cap(Q2))

	summer := months[6:9]
	//summer长度为item个数9-6=3
	fmt.Println("summer's len is:", len(summer))
	//summer容量为len(months)减去months从下标6开始截取的位置为13-6=7
	fmt.Println("summer's cap is:", cap(summer))
	fmt.Println("---------------------")
	//由于Q2、summer，底层数组都是months，并且他们三个都共有{6：“liuyue”}这块空间的引用，这块空间的值一旦被改就全改
	Q2[2] = "Q2 -> 6"
	fmt.Println("Q2[2]:", Q2[2])
	fmt.Println("months[6]", months[6])
	fmt.Println("summer[0]", summer[0])
	fmt.Println("---------------------")
	//这次summer来改
	summer[0] = "summer -> 6"
	fmt.Println("Q2[2]:", Q2[2])
	fmt.Println("months[6]", months[6])
	fmt.Println("summer[0]", summer[0])
	fmt.Println("---------------------")
	//最后由months改回去
	months[6] = "liuyue"
	fmt.Println("Q2[2]:", Q2[2])
	fmt.Println("months[6]", months[6])
	fmt.Println("summer[0]", summer[0])

	fmt.Println("---------------------")
	//新定义Q2Q3，从Q2的第一个元素开始，往后截取6个，虽然超过了len(Q2),但由于仍在底层数组months的地址范围，等价于months[4:10]的操作
	Q2Q3 := Q2[:6]
	fmt.Println("Q2Q3:", Q2Q3)
	//Q2Q3长度为10-4=3
	fmt.Println("Q2Q3's len is:", len(Q2Q3))
	//summer容量为len(months)减去months从下标4开始截取的位置为13-4=9
	fmt.Println("Q2Q3's cap is:", cap(Q2Q3))

	fmt.Println("---------------------")
	//slice的声明方法：
	//方法一、直接声明一个len(s) = cap(s)，反正可以重新赋值
	s1 := []int{99: 99}
	fmt.Println(s1)
	fmt.Println("s1's len is:", len(s1))
	fmt.Println("s1's cap is:", cap(s1))
	//想什么时候赋值都行
	for k, v := range s1 {
		if k != 0 && v == 0 {
			s1[k] = k
		}
	}
	fmt.Println(s1)

	fmt.Println("---------------------")
	//方法二、原理同上，最后通过s[:]全部引用数组s（或者先引用再赋值）
	s := [...]int{99: 99}
	for k, v := range s {
		if k != 0 && v == 0 {
			s[k] = k
		}
	}
	fmt.Println(s[:])

	fmt.Println("---------------------")

	//方法三、使用make函数
	//make([]T, len)
	//make([]T, len, capability) // same as make([]T capability)[:len]
	s3 := make([]int, 100, 100)
	//赋值
	s3 = s[:]
	fmt.Println(s3)

	fmt.Println("---------------------")

	a := []int{1, 2, 3}

	fmt.Println("a's len is:", len(a))
	fmt.Println("a's cap is:", cap(a))

	a = append(a, 4, 5, 6)
	fmt.Println("a's len is:", len(a))
	fmt.Println("a's cap is:", cap(a))
	fmt.Println(a)

	fmt.Println("---------------------")
	b := new([]int)
	*b = append(*b, 1, 2, 3, 4, 5, 6)
	fmt.Println(*b)
}
