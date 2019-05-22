package main

import (
	"fmt"
	"time"
)

//32.timers  定时器？计时器？
func main()  {
	timer1:=time.NewTimer(2*time.Second)   //创建一个两秒的计时器
	<-timer1.C  //go为每个定时器提供了通道，定时器过期后，它会发送一个指示计时器过期的值
	fmt.Println("Timer 1 expired")

	timer2:=time.NewTimer(time.Second)  //创建一个1秒的定时器

	<-timer2.C
	fmt.Println("Time 2 expired1")

	go func() {
		<-timer2.C
		fmt.Println("Time 2 expired")
	}()

	stop2:=timer2.Stop()    //这句是什么呢？就是说我们可以在定时器自己过期之前就给它手动停止
	                        //而我们既然手动停止定时器了，那么它自然也就不会打印过期语句了
	if stop2{   //当定时器停止时，返回true
		fmt.Println("Timer 2 stopped")
	}
}
//总结32：这个嘛，记住各个函数的作用就行了，其他没什么
