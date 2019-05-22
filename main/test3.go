package main

import (
	"fmt"
)

func main()  {
	fmt.Println(plus(1,2))
	fmt.Println(plusThree(1,2,3))

	fmt.Println(mulivalues())
	a,b:=mulivalues()
	fmt.Println(a)
	fmt.Println(b)

	sum(1,2,3)
	nums:=[]int{1,2,3}
	sum(nums...)    //这三个点真的很有灵性
	//注：go语言不允许混合不同类型的元素

	nextInt:=intSeq()
	fmt.Println(nextInt())  //1
	fmt.Println(nextInt())  //2
	fmt.Println(nextInt())  //3
	//这个i是静态的不成？

	newInt:=intSeq()
	fmt.Println(newInt())   //1
	//观察这写返回值，表现了闭包的记忆效应
	//被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随着闭包生命期一直存在
	//闭包本身就如同变量一样拥有了记忆效应
	//目前还不知道闭包的具体应用场景

	fmt.Println(fact(7))

}

//12.函数
func plus(a int,b int)int{   //go的函数创建与python无异，但有返回值要标注返回值
	return a+b
}
func plusThree(a,b,c int)int{
	return a+b+c
}

//13.函数，多返回值
func mulivalues() (int,int) {
	return 3,7
}
//总结13：从这点来看，go的优势明显，只要在函数体标定返回值即可实现多返回

//14.动态参数函数
func sum(nums...int)  {
	total:=0
	for _,v:=range nums{
		total+=v
	}
	fmt.Println(total)
}
//总结14：go语言的优点慢慢体现出来了

//15.闭包 函数+引用环境=闭包
func intSeq() func() int {  //函数名：intSeq，返回值：func()（匿名函数）,int为匿名函数的返回值
	i:=0
	return func() int {
		i++
		return i
	}
}
//感觉go语言关于返回值的标定是很重要的

//16.递归
func fact(i int)int  {
	if(i==0){
		return 0
	}else {
		return fact(i-1)
	}
}
//总结16：递归嘛，大同小异，注意跳出条件就可以了


