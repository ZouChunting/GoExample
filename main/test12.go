package main

import "fmt"

//26.通道方向
//channel这部分的确是go语言的重点部分，也是难于理解的部分
//本来看一遍就能手撸代码，现在只能先根据代码理解意思
//到目前为止，go语言难以理解的部分在指针和协程这两部分

//当使用通道作为函数参数时，可以指定通道是否仅用于发送或接收值。这种特异性增加了程序的类型安全
func ping(pings chan <-string,msg string)  {  //参数：写值通道，string
	pings<-msg  //将msg的值写入通道
	//value:=<-pings   //基于pings的写值特性，我们如何想从其读值，会报错，大意就是只能用于写值不能读值的意思
}//ping函数只接受用于写值的通道

func pong(pings<-chan string,pongs chan <- string)  {  //参数：读值通道，写值通道
	msg:=<-pings  //定义变量，用于接受读值通道pings中的值
	pongs<-msg   //将变量中的值写到写值通道pongs中
	//pongs<-pings  //本来想着为啥偏要有个中介msg，
	//这种写法的后果是会报错，can not use 'pings'(type-chan string) as type string in send
}//pong函数只接受用于读值的通道

func main()  {
	pings:=make(chan string,1)  //创建类型为string，缓存为1的通道
	pongs:=make(chan string,1)  //俺也一样
	ping(pings,"passed message")  //调用ping函数
	pong(pings,pongs)
	fmt.Println(<-pongs)
}
//总结26：不太明白这个为啥要单独摘出来讲，不过看到这，已经能明显感受到go语言的独特之处
