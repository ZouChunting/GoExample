package main

import (
	"fmt"
	"time"
)

//28.超时

func main()  {
	//创建通道
	c1:=make(chan string)
	//创建协程
	go func() {
		time.Sleep(2*time.Second)
		c1<-"result 1"
	}()

	select {
	case res:=<-c1:
		fmt.Println(res)
	case <-time.After(time.Second):   //go里的函数都有点迷，但又蜜汁人性化
		fmt.Println("timeout 1")
	}
	//这里值得注意的一点是，两个case都执行了，但由于上面的协程需要两秒以上才能完成赋值，所以c1中仍没有值，因此没有输出

	c2:=make(chan string)
	go func() {
		time.Sleep(2*time.Second)
		c2<-"result 2"
	}()
	select {
	case res:=<-c2:
		fmt.Println(res)
	case <-time.After(3*time.Second):  //<-time.After(3*time.Second)样子有点怪异，但是稍微记忆一下就好
		fmt.Println("timeout 2")
	}
	//这里的两个case也是都执行了的，但是协程在3秒内完成赋值，因此不满足第二个case的条件，因此没有打印
	//之所以说这么多，只是想多表现switch和select的区别

	//总结28：了解了select在通道中的作用，以及几个time函数，这边很好理解
}
