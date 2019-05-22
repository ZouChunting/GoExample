package main

import "fmt"

func main()  {
	i:=1
	fmt.Println("initial:",i)
	zeroval(i)
	fmt.Println("zeroval:",i)
	zeroptr(&i)  //通过指针赋值
	fmt.Println("zeroptr:",i)
	fmt.Println(&i)

	//go语言指针
	//在对普通变量使用&操作符取地址获得这个变量的指针后，可以对指针使用*操作，即指针取值
	//指针我学的也不太好，先了解一下，以后根据实例理解吧
}

//17.指针，学c时谈虎色变的指针
func zeroval(ival int)  {
	ival=0
}
func zeroptr(iptr *int)  {
	*iptr=0
}