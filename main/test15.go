package main

import "fmt"

//29.非阻塞通道操作

//Non-Blocking Channel Operations非阻塞通道操作，不知道这样翻译对不对
//感觉这个概念不是很好理解
//通道上的基本发送和接收是阻塞的。
//但是，我们可以使用含有default子句的select来实现非阻塞发送，接收。
func main()  {
	messages:=make(chan string)
	signals:=make(chan bool)

	select {
	case msg:=<-messages:
		fmt.Println("received message:",msg)
	default:
		fmt.Println("no message received")
	}
	//我们先来理解一下这段的select-default
	//以上我们已经创建了通道messages，但由于message还是nil，因此被赋值的msg也是nil，也就是说
	//select没有等到msg:=<-messages的赋值操作，就立即采取了default方案，并且结束此次select
	//更官方的说法是：这是一个非阻塞接收。

	msg:="hi"
	select {
	case messages<-msg:
		fmt.Println("sent messgae:",msg)
	default:
		fmt.Println("no message sent")
	}
	//非阻塞发送的工作方式类似。这里的msg不能发送到messages通道，因为通道没有缓冲区，也没有接收器。因此选择default方案。
	//这里需要说明的是，什么时候msg可以发送到messages通道里，简单来说，要么通道定义时设置了缓冲区
	//要么创建了一个协程去发送，否则会报错
	//通道是go语言里比较晦涩的概念，需要更多的例子去理解

	select {
	case msg:=<-messages:
		fmt.Println("received message",msg)
	case sig:=<-signals:
		fmt.Println("received signal",sig)
	default:
		fmt.Println("no activity")
	}
	//多路非阻塞选择。在这里，我们试图非阻塞接收messages和signals。
	//这里的前两个case为什么没被执行，我想已经显而易见了。messages和signals一直都是空的。
}
//补充：我发现前面对于switch和select的解释太浅显了，更详尽的说法是
//（1）Go会依照从上至下的顺序对每一条case语句中case表达式进行求值，
//只要被发现其表达式与switch表达式的结果相同，该case语句就会被选中。其余的case语句会被忽略。
//（2）select 的 case 里的操作语句只能是【IO 操作】
//select 会一直等待等到某个case语句完成才结束。
//break语句也可以被包含在select语句中的case语句中。
//它的作用是立即结束当前的select语句的执行。不论其所属的case语句中是否还有未被执行的语句。

//纵然如此，我其实还不是很确定前面关于switch和select的说法对不对
