package main

import (
	"fmt"
	"math"
)

//20.接口
type geometry interface {
	area1() float64
	perium1() float64
} //抽象方法

type rectangle struct {
	width,height float64
}//方形结构体
type circle struct {
	radius float64
} //圆形结构体

func (r rectangle) area1() float64  {
	return r.height*r.width
}//计算方形面积
func (r rectangle) perium1() float64 {
	return 2*r.width+2*r.height
}//计算方形周长

func (c circle) area1() float64 {
	return math.Pi*c.radius*c.radius
}//计算圆形面积
func (c circle) perium1() float64 {
	return 2*math.Pi*c.radius
}//计算圆形周长

func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area1())
	fmt.Println(g.perium1())
}

func main()  {
	r:=rectangle{3,4}
	c:=circle{5}

	measure(r)
	measure(c)
}
//总结20：我们可以先回忆一下java的接口是怎样实现的，方形类和圆形类都有面积和周长方法，但是计算方法不同，我们把这两个方法
//抽象到一个接口里，方形类和圆形类implements这个接口来实现各自的两种方法
//go语言的接口也可以这么理解，"一个对外接口，多个内在实现方法"，多态
