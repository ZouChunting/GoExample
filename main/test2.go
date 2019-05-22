package main

import (
	"fmt"
)

func main(){
	//8.数组
	var a[5]int   //定义长度为5的int类型一维数组
	fmt.Println("emp:",a)     //默认值为0
	a[4]=100      //数组为指定index赋值
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])

	fmt.Println(len(a))   //基本操作

	b:=[5]int{1,2,3,4}   //创建数组后，并不一定要把值赋全，这一点似乎比java方便一点，java的变量创建的确是大部分语言里最繁琐的
	fmt.Println(b)

	var twoD [2][3]int  //定义二维数组
	for i:=0;i<len(twoD);i++{
		for j:=0;j<len(twoD[0]);j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d:",twoD)
	//总结8：Go的数组和python如出一辙，向人类能够轻易理解的方向走
	//但是越是人类能够理解的高级语言，往往在底层的工作性能越低，不知道Go的性能怎样

	//9.切片
	//首先，为什么要使用slice，因为在go中，数组是不可变的
	//java中有arraylist等支持可变长度的数组，stringbuilder等支持可变长度的字符串
	//go通过切片来实现数组长度变化，更加灵活地运用数组。

	//slice声明，使用make函数，创建形式与数组创建类似
	s:=make([]string,3)    //类型为int，长度3，可添加容量
	fmt.Println("emp:",s)
	//slice 的声明也可以像 array 一样，只是不需要长度 s:=[]string

	s[0]="aaa"
	s[1]="bbb"
	s[2]="ccc"
	fmt.Println("set:",s)   //赋值，到这里还没有看出slice的特别之处
	fmt.Println("len:",len(s))

	//追加元素
	s= append(s, "ddd","eee","fff")

	fmt.Println("after append:",s)

	//拷贝
	c:=make([]string,len(s))
	copy(c,s)  //将s的内容拷贝到c，注意函数两元素的顺序
	fmt.Println("copy:",c)

	//开始切片
	l:=s[2:5]
	fmt.Println("slice s to l:",l)  //slice s to l: [ccc ddd eee]
	l=s[2:]
	fmt.Println("slice s to l:",l)  //slice s to l: [ccc ddd eee fff]
	l=s[:5]
	fmt.Println("slice s to l:",l)  //slice s to l: [aaa bbb ccc ddd eee]
	//怎么说呢，熟悉python的人看到这种所谓的切片肯定很熟悉，貌似和python的确没啥区别，除了名字高端一点

	//二维数组切片
	twoD2:=make([][]int,3)
	//稍微解释下下面这段
	for i:=0;i<3;i++{  //行循环
		length:=i+1    //取各行列数为i+1
		twoD2[i]=make([]int,length)   //每行以length创建切片
		for j:=0;j<length;j++{
			twoD2[i][j]=i+j   //普通的赋值操作
		}
	}
	fmt.Println("twoD2:",twoD2)
	//总结9：切片照着python理解即可

	//10.map
	//创建，通过make函数
	m:=make(map[string]int)  //string是key，int是value
	m["k1"]=7  //简单的赋值操作
	m["k2"]=13
	fmt.Println("set:",m)

	v1:=m["k1"]
	fmt.Println("v1:",v1)  //取出某value

	fmt.Println("len:",len(m))  //长度

	delete(m,"k2")  //以key删除某元素
	fmt.Println("after delete:",m)

	//判断key是否存在
	_,prs:=m["k1"]
	fmt.Println("prs:",prs)  //true
	_,prs=m["k2"]
	fmt.Println("prs:",prs)  //false

	if v, ok := m["a"]; ok {   //不得不说，这步操作有点6
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")
	}

	//遍历map
	for k:=range m{
		fmt.Println(k)
	}
	for _,v:=range m{
		fmt.Println(v)
	}
	for k,v:=range m{
		fmt.Println(k,v)
	}

	n:=map[string]int{"aaa":1,"bbb":2}  //创建时指定类型赋值则不需要用make
	fmt.Println(n)
	//总结10：map就是一个key-value集，java中的hashtable有相似作用，python中也有map，dict

	//11.range
	num:=[]int{1,2,3}
	sum:=0
	for _,num:=range num{  //_,num即index,value
		sum+=num
	}
	fmt.Println(sum)

	for index,value:=range num{
		if value==2{
			fmt.Println(index)
		}
	}

	//举一反三，上面map的遍历就是这么来的
	for i,c:=range "go"{    //0 103   index,ascii码
		fmt.Println(i,c)    //1 111
	}

}
