package main

import (
	"fmt"
)


func main()  {
	//23.channel
	//什么是channel呢？通道的意思
	//channel是goroutine之间通信的一种方式
	//channel提供了一种通信机制，通过它，一个goroutine可以向另一goroutine发送消息。
	//channel本身还需要关联一个类型，也就是channel可以发送数据的类型。
	//例如，发送int类型消息的channel写作chan int

	//channel的基础操作
	//创建channel
	//ch:=make(chan int)
	//channel的读写操作
	//ch<-11  //将int类型的11写入channel
	//ch<-12
	//q<-ch或者q=<-ch  读取channel
	//关闭channel
	//close(ch)

	//创建channel
	messages:=make(chan string)

	//启动一个线程向通道里写信息
	go func() {
		messages<-"ping"
	}()

	//将通道里的数据读取出来
	msg:=<-messages
	fmt.Println(msg)
	//总结23：理解channel不是一件难事，要将它融入具体的应用场景，这里，我只熟悉概念，以后有机会再拓展

	//24.有缓存的通道
	//有缓存的channel的声明方式为指定make函数的第二个函数，该参数为channel缓存的容量
	//ch:=make(chan int,10)
	//有缓存的channel类似一个阻塞队列。当缓存未满时，向channel中发送消息时不会阻塞，当缓存满时，发送操作将被阻塞
	//直到有其他goroutine从中读取消息；相应地，当channel中消息不为空时，读取消息不会出现阻塞，当channel为空时，读取操作会
	//造成阻塞，直到有goroutine向channel中写入消息

	c:=make(chan int,2)
	c<-1
	c<-2
	//c<-3  缓存为2，当下没有被唤醒的进程，因此若消息超过2会报错
	fmt.Println(<-c)
	fmt.Println(<-c)
	c<-3
	fmt.Println(<-c)  //前面已有读取操作，因此再次写入数据不会报错
}



