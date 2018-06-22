package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

//判断map的值是否相等
func mapEqual(m1 map[string]int, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	//一定要判断!ok，否则当m1["xhtian"]因为不存在，返回零值0,而正好比较函数m2["xhtian"] = 0时，函数就会错误判断两个map相等
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

//读取文件中的行，模拟一个string list,去重
func newStrList(fileName string) map[string]bool {
	m := make(map[string]bool)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil
	}
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\r\n") {
		if !m[line] {
			m[line] = true
		}
	}
	return m
}

func main() {

	//字面量声明
	m1 := map[string]int{
		"kook":  36,
		"ping":  34,
		"chang": 8,
	}
	fmt.Println(m1)
	fmt.Println("1---------------------")
	//make函数声明
	m2 := make(map[string]int)
	m2["kook"] = 36
	m2["ping"] = 34
	m2["chang"] = 8
	fmt.Println(m2)
	fmt.Println("2---------------------")
	//遍历map
	for k, v := range m1 {
		fmt.Printf("the key is:%s,the value is:%d\n", k, v)
	}
	fmt.Println("3---------------------")
	//删除map的指定key，即使不存在也无妨
	delete(m1, "noname")
	fmt.Println(m1)
	fmt.Println("4---------------------")
	//按照key进行排序,先make一个slice
	names := make([]string, 0, len(m1))
	for name, _ := range m1 {
		names = append(names, name)
	}
	//调用sort函数，对slice数据排序
	sort.Strings(names)

	//排序后输出
	for _, name := range names {
		fmt.Printf("%s's age is:%d\n", name, m1[name])
	}
	fmt.Println("5---------------------")
	//map中，当值不存在，则返回数据类型的零值，那如何判断map中元素是否存在呢？
	if _, ok := m1["tian"]; !ok {
		fmt.Println("m1['tian'] is null")
	}
	fmt.Println("6---------------------")
	var m3 map[string]int
	fmt.Println("m3 is null?", m3 == nil)
	fmt.Println("m3's len?", len(m3))

	m4 := map[string]int{}
	fmt.Println("m4 is null?", m4 == nil)
	fmt.Println("m4's len?", len(m4))

	m5 := make(map[string]int)
	fmt.Println("m5 is null?", m5 == nil)
	fmt.Println("m5's len?", len(m5))
	fmt.Println("7---------------------")
	//判断两个map是否相等。
	b := mapEqual(map[string]int{"A": 42}, map[string]int{"B": 42})
	fmt.Println("mapEqual func invork... and return [", b, "]")
	fmt.Println("8---------------------")
	sm := newStrList("testfile.txt")
	fmt.Println(sm)
}
