package main

import "fmt"

//19.方法 结构体加方法就是完全体的java class吧
//go中任何自定义类型都可以有方法，不仅仅是struct
type rect struct {
	width,height int
} //方形结构体

func (r rect) area() int  {  //对自定义类型rect定义一个方法 area
	return r.height*r.height
}//计算面积

func (r *rect) perim() int {
	return r.width*2+r.height*2
}//计算周长

func main()  {
	r:=rect{4,5}
	fmt.Println("面积",r.area())
	fmt.Println("周长",r.perim())
	rp:=&r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
	//我发现所有概念一和指针扯在一起，理解难度就上一个level，指针真是我永远的痛
}
