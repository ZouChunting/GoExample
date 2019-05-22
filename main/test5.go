package main

import "fmt"

//18.结构体 就相当于java的class吧
type person struct {
	name string
	age int
} //稍微注意下创建格式就好

func main()  {
	p1:=person{"aa",12}
	fmt.Println(p1)  //{aa 12}
	fmt.Println(&p1)  //&{aa 12}
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	p2:=&p1
	fmt.Println(p2.age)
	p2.age=22
	fmt.Println(p2.age)
	fmt.Println(p1.age)
	//加了指针，我真的会有点乱，但这个例子不是很复杂，能理解
}
