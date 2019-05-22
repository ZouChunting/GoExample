package main

import (
	"fmt"
	"time"
)

//35.速率限制
//通过滴答器控制响应速度
//速率限制是控制资源利用和保持服务质量的重要机制。

func main()  {
	//---a---
	requests:=make(chan int,5)
	for i:=1;i<=5;i++{
		requests<-i
	}
	close(requests)
	//创建缓冲通道，赋值后关闭

	limiter:=time.Tick(200*time.Millisecond)  //设置200ms的滴答器

	for req:=range requests{
		<-limiter
		fmt.Println("request",req,time.Now())
	}//实现每200ms接收一次的效果

	//咦，突然发现了关闭通道的作用
	//首先，通道不能给gc自动回收（我觉得我这话说的很外行，gc我懂的也不是很多）
	//emmm怎么解释呢？简单来说，如果我们不手动关闭通道，那么通道就是一直存在的
	//上面的range通道语句就会一直遍历，通道存在，遍历就不会停
	//那么，问题又来了，close以后就宣布的通道的死亡吗？能不能重新打开或者启动通道。
	//如果将close(requests)注释掉，那么程序就迟迟不能结束，停在
	//request 5 2019-05-16 13:33:22.656225 +0800 CST m=+1.005354459
	//通道的关闭是在告诉接收者不会再提供新的值
	//for-range语句会自动检测通道是否关闭

	//---a---
	fmt.Println("")

	//---b---
	burstyLimiter:=make(chan time.Time,3)
	for i:=0;i<3;i++{
		burstyLimiter<-time.Now()
	}//创建缓冲通道并赋值

	go func() {
		for t:=range time.Tick(2*time.Second){    //200*time.Millisecond
			burstyLimiter<-t
		}
	}()//在协程中向burstyLimiter写值，每200ms写一次

	burstyRequests:=make(chan int,5)
	for i:=1;i<=5;i++{
		burstyRequests<-i
	}
	close(burstyRequests)
	//创建通道，赋值后关闭

	for req:=range burstyRequests{
		<-burstyLimiter
		fmt.Println("request",req,time.Now())
	}
	//我们遍历的通道burstyRequests里有5个值，因此会有5个输出结果
	//遍历burstyRequests的输出速率是由burstyLimiter决定的
	//burstyLimiter叫做限制器，当我们能从其中接收到值时（可以理解为限制信号）burstyRequests才可以打印输出
	//如果对限制器的概念不是很了解，可以把滴答器的时间修改的长一点，就可以发现前3次的打印很迅速，因为通道中
	//有初始赋的3个值
	//后两次的输出较慢，因为限制器中的通道写值放到了协程中，限制时间写入
	//---b---
}
