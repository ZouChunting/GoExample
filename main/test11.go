package main

import (
	"fmt"
	"time"
)

//25.信道同步
//下面这段代码是想实现什么呢？
//这是一个使用缓存通道（阻塞队列）来等待goroutine完成的例子
func worker(done chan bool)  {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done<-true  //将true写入channel
}

func main()  {
	done:=make(chan bool,1)  //创建缓存为1，类型为bool的信道
	go worker(done)

	fmt.Println(<-done)  //或者这里也可以直接写<-done，只不过不打印值而已
	//不加这行，也就是说，不特地接收线程执行的值，这程序就不管goroutine是否执行完毕，主程序就直接退出了
	//这，emmm，算是语言特色吧
	//众所周知，我们在java中随便写一个线程，继承thread类，start线程，不等到run执行完毕，主程序是不会exit的
	//emmm，区别先写在这，以后再研究原因吧
}
