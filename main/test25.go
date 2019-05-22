package main

import (
	"fmt"
	"sort"
)

//39.排序
func main()  {
	strs:=[]string{"c","a","b"}
	sort.Strings(strs)
	fmt.Println("Strings:",strs)
	//基本所有高级语言里都有sort排序，python和go调用起来都比较简单，java需要实现排序接口

	ints:=[]int{7,2,4}
	sort.Ints(ints)
	fmt.Println("Ints:",ints)

	s:=sort.IntsAreSorted(ints)
	fmt.Println("Sorted:",s)
	//这个就比较有意思了，可以判定是否排序过

	ints2:=[]int{1,2,3}
	s2:=sort.IntsAreSorted(ints2)
	fmt.Println("Sorted2:",s2)
	//比较一下这两段，也就是说，即便没有调用过sort，但是有序的，返回的仍是true
}
//总结39：排序的概念算是老生常谈了，不知道到各大排序算法在web项目中的应用几何，但我认为还是很有必要
//了解各大排序算法的特点和使用场景的，尤其是计数排序，我觉得特神奇
