package main

import (
	"fmt"
	"sort"
)

//40.排序函数
//下面这段很明显是根据字符串长度进行排序
type bylength []string
//emmm，这是定义了一个string数组类型的结构体
//这里就比较能觉出java的规范了，java语言虽然有时较为繁琐，但是代码是美观整洁规范的

func (s bylength) Len()  int {
	return len(s)
}//返回数组长度--定制排序接口1

func (s bylength) Swap(i,j int)  {
	s[i],s[j]=s[j],s[i]
}//交换，这里go和python是一样的，而java需要一个tmp来做交换中介值，至于为啥，细枝末节的，我也不想深究了
//--定制排序接口2

func (s bylength) Less(i,j int) bool  {
	return len(s[i])<len(s[j])
}//返回两个下标的长度比较结果--定制排序接口3
//这里可以理解为排序标准
//其实java中也有类似的，就是返回0,-1,1的那个，学名我忘了叫啥了

func main()  {
	fruits:=[]string{"peach","banana","kiwi"}
	sort.Sort(bylength(fruits))
	fmt.Println("Sorted:",fruits)
}
//总结40：这里值得一提的是go的定制排序
//有时候，我们想要按照一个集合的自然顺序以外的东西来排序。例如，我们想按字符串的长度排序，而不是按字母顺序排序。
//首先要创建自定义类型type bylength []string，
//其次，在此类型上实现3个接口
//func (s bylength) Len()  int、func (s bylength) Swap(i,j int)、func (s bylength) Less(i,j int) bool
//然后调用sort就可以了
//实现的三个接口，实现了返回数组长度、按下标交换（确定排序操作）、确定排序标准3个功能
//这个多写几个例子就熟悉了，以后有应用场景再研究吧