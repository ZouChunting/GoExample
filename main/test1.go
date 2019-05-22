package main

import (
	"fmt"
	"math"
	"time"
)

//for循环
//var arr=[...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}  //创建数组

func main()  {
	//1.打印hello world
	fmt.Println("hello world")
	/*
	fmt.Println(arr[0])//打印arr[0]
	//循环遍历1
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}
	//循环遍历2

	//循环遍历3
	 */
	//2.变量基本类型与逻辑运算
	fmt.Println("1+1=",1+1)  //整型变量
	fmt.Println("go"+"land")  //字符型变量
	fmt.Println("7.0/3.0=",7.0/3.0)   //浮点数
	fmt.Println(true&&false)  //布尔型  与
	fmt.Println(true||false)  //或
	fmt.Println(!true)    //非
	//总结2：变量基本类型与java基本无异

	//3.自定义变量
	var a = "initial"  //字符串类型
	fmt.Println(a)

	var b,c int = 1,2  //整型
	fmt.Println(b,c)

	var d = true   //布尔类型
	fmt.Println(d)

	var e int     //指定类型不赋值，整型默认值为0
	fmt.Println(e)
	e=3     //给已定义变量赋值
	fmt.Println(e)

	f:="short"    //简易赋值，只可以是未定义变量
	fmt.Println(f)
	//总结3：go中自定义变量，赋值或指定类型，必有其一，否则不合法

	//4.数学运算
	const s string  = "constant"    //静态变量，不可修改，即不可另行赋值，也并不需要一定被引用
	fmt.Println(s)

	const n  =  500000000  //5亿
	const m  = 3e20/n
	fmt.Println(m)    //6e+11
	fmt.Println(int64(m))   //600000000000     64进制转10进制

	fmt.Println(math.Sin(n))    //数学运算
	//总结4：与java基本无异，java静态变量static值是可更改的，go的const不给更改赋值

	//5.for
	i:=1
	for i<3{        //此处的for与java中的while功能一致
		fmt.Println(i)
		i++
	}

	for j:=7;j<9;j++{    //与java中普通的for循环无异
		fmt.Println(j)
	}

	for{
		fmt.Println("loop")
		break   //循环一次跳出
	}

	for k:=0;k<5;k++{
		if k%2==0{
			continue
		}
		fmt.Println(k)
	}

	//6.if...else...
	if 7%2==0{
		fmt.Println("7 is even")   //偶数
	}else {
		fmt.Println("7 is odd")   //奇数
	}
	if num:=-1;num<0{     //在if中定义变量，并以此作为判断分支
		fmt.Println(num,"num<0")
	}else if num<10 {
		fmt.Println("num<10")
	}else {
		fmt.Println("else")
	}

	const v1  =  2
	if v1>0&&v1==2{     //带与判断
		fmt.Println("v1>0&&v1==2")
	}else {
		fmt.Println("!v2")
	}
	//总结6：与java无异

	//7.switch
	v:=1
	switch v {
	case 1:
		fmt.Println("v=1")
	case 2:
		fmt.Println("v=2")
	default:
		fmt.Println("v=0")
	}

	switch time.Now().Weekday() {  //这个取时间的函数有点意思
	case time.Saturday,time.Sunday:
		fmt.Println("放假")
	default:
		fmt.Println("工作")
	}

	t:=time.Now()
	switch  {     //switch中不一定要有条件
	case t.Hour()<12:
		fmt.Println("上午")
	case v==1:
		fmt.Println("v=1")
	case t.Hour()>12:
		fmt.Println("下午")  
	}//分支执行不超过一条，java中除非在分支中加break，否则每条分支语句都会向下对比

	whatAmI:= func(in interface{}) {
		switch t:=in.(type){
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("nowhere man",t)
		}
	}
	whatAmI(12)
	whatAmI(true)
	whatAmI("qwer")    //还没有看到interface，这里后续再详看
	//总结7：基本的switch语句需要注意go的语法习惯
}






