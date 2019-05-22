package main

import (
	"fmt"
	"time"
)

//27.select
func main()  {
	c1:=make(chan string)  //创建通道
	c2:=make(chan string)

	//创建协程1，赋值操作
	go func() {
		time.Sleep(time.Second)
		c1<-"one"
	}()
	//创建协程2，赋值操作
	go func() {
		time.Sleep(time.Second)
		c2<-"two"
	}()
	//等待1，2协程的操作结果
	for i:=0;i<2;i++{
		select {  //虽然这里有case，但是select和switch不一样，前面写过switch，它只会执行一条及以下分支
		case msg1:=<-c1:
			fmt.Println("received:",msg1)
		case msg2:=<-c2:
			fmt.Println("recevied:",msg2)
		}
	}
}
//总结27：我们通过select实现什么呢？等待接收多个通道的操作
//将协程、通道、选择相结合是go的一个强大特性
