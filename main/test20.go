package main

import (
	"fmt"
	"time"
)

//34.工作池
//这个例子展示了如何使用协程和通道实现工作池
func Worker(id int,jobs<-chan int,results chan<- int)  {  //参数：工人编号，读值通道，写值通道
	for j:=range jobs{
		fmt.Println("worker",id,"started job",j)  //工人开始一件工作的标记
		time.Sleep(time.Second)
		fmt.Println("worker",id,"finished job",j)  //工人完成一件工作的标记
		results<-j*2    //将工作完成标记存入结果通道
	}
}

func main()  {
	fmt.Println(time.Now())
	jobs:=make(chan int,100)
	results:=make(chan int,100)

	for w:=1;w<=3;w++{
		go Worker(w,jobs,results)
	}//有3个工人，3个工人并发工作
	//此时有3个协程在并发工作

	for j:=1;j<=5;j++{
		jobs<-j
	}//写值，共有5件工作
	close(jobs)

	for a:=1;a<=5;a++{
		<-results
	}//收集5件工作完成的结果，聪明的小机灵鬼应该知道为什么要收集，因为不收集，协程就执行不完啦
	fmt.Println(time.Now())

	//观察一下时间标记，一个工人完成一件工作的时间是1秒，然而此程序执行时间不超过3秒，这就是并发的好处了
	//高并发是go的优势性能之一
}
