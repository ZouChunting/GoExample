package main

import "fmt"

//31.通道的遍历

func main()  {
	//创建一个缓冲通道
	queue:=make(chan string,2)
	queue<-"one"
	queue<-"two"
	close(queue)
	for elem:=range queue{
		fmt.Println(elem)
	}
}
//总结31：这个比较好理解了
