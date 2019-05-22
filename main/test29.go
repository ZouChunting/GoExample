package main

import (
	"fmt"
	"strings"
)

//43.集合函数
//我就纳了闷了，演示什么需要写这么长
//go语言不支持泛型
//我之前学java的时候，只知道有这么个概念，现在来补补课，什么是java泛型
//java泛型提供了编译时类型安全检测机制，该机制允许程序员在编译时检测到非法的类型。
//泛型的本质是参数化类型，也就是说所操作的数据类型被指定为一个参数
//使用java泛型的概念，我们可以写一个范型方法来对一个对象数组排序。
//然后，调用该泛型方法来对整型数组、浮点数组、字符串数组等进行排序

func Index(vs []string,t string) int {
	for i,v:=range vs{
		if v==t{
			return i
		}
	}
	return -1
}//此函数返回目标字符串在字符串数组中下标，若没有找到，则返回-1

func Include(vs []string,t string) bool  {
	return Index(vs,t)>=0
}//此函数返回目标字符串是否被包含在字符串数组中，true/false，
//此函数调用了Index，因为下标为-1就是没找到，不存在嘛

//以上两个函数其实没啥幺蛾子，主要看看下面几个

func Any(vs []string,f func(string) bool) bool {
	for _,v:=range vs{
		if f(v){
			return true
		}
	}
	return false
}//这种将函数作为参数传入函数的，之前不是没见过，但看见了，就感觉go与java比起来，就不算太规范
//难道这就是传说中、大名鼎鼎的集合函数嘛

func All(vs []string,f func(string) bool) bool {
	for _,v:=range vs{
		if !f(v){
			return false
		}
	}
	return true
}//参数：字符串数组，参数字符串、返回布尔值的函数
//整个函数的意思就是，遍历字符串数组，将每个字符串传入f函数，根据函数结果，决定返回值
//但是这些整个函数到底起什么作用，还是以传入的函数参数决定
func Filter(vs []string,f func(string) bool) []string {
	vsf:=make([]string,0)
	for _,v:=range vs{
		if f(v){
			vsf=append(vsf,v)
		}
	}
	return vsf
}

func Map(vs []string,f func(string) string) []string {
	vsm:=make([]string,len(vs))
	for i,v:=range vs{
		vsm[i]=f(v)
	}
	return vsm
}

func main()  {
	var strs=[]string{"peach","apple","pear","plum"}

	fmt.Println(Index(strs,"pear"))  //返回"pear"在数组中的下标
	fmt.Println(Include(strs,"grape"))  //数组中是否存在"grape"

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v,"p")
	}))
	//strings.HasPrefix ???
	//既然如此，我们看一下strings包的用法吧
	//s:="hello,world!!!!"
	//a:=strings.Count(s,"!")
	//fmt.Println(a)  //4
	//b:=strings.Contains(s,"!!")
	//fmt.Println(b)   //true
	//c:=strings.ContainsAny(s,"abcde")
	//fmt.Println(c)   //true
	//e:=strings.Index(s,"l")
	//fmt.Println(e)    //2
	//f:=strings.LastIndex(s,"l")
	//fmt.Println(f)   //9
	//g:=strings.HasPrefix(s,"hello")
	//fmt.Println(g)   //true
	//h:=strings.HasSuffix(s,"!")
	//fmt.Println(h)   //true
	//i:=strings.ToUpper(s)
	//fmt.Println(i)   //HELLO,WORLD!!!!
	//试试这些代码，你将明白一切，那么strings.HasPrefix(v,"p")的意思就是v字符串是否以p开头

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v,"p")
	}))//返回true/false，判断strs是否每个字符串都是由p开头

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v,"e")
	}))//返回了一个字符串数组，为str字符串数组中所欲包含e的字符串

	fmt.Println(Map(strs,strings.ToUpper))
	//将字符串数组中所有字符串数组变为大写并返回字符串数组
	//strings.ToUpper竟然是可以直接作为函数参数的，意料之外，情理之中
	fmt.Println(Map(strs,strings.ToLower)) //那么我也可以这样写喽
}
//总结43：看完这节，我感到的还是浓浓go特色的自定义，这和泛型有什么关系？？？小小的脑袋，大大的疑惑
