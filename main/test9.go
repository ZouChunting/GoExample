package main

import "fmt"

//22.goroutines
//相当于java的线程
//goroutine是建立在线程之上的轻量级抽象。它允许我们以非常低的代价在同一个地址空间中并行地执行多个函数
//或者方法。相比于线程，它的创建和销毁的代价要小很多，并且它的调度是独立于线程的。使用go关键字即可创建。
func f(from string)  {   //参数：自定义变量from，类型string
	for i:=0;i<3;i++{
		fmt.Println(from,":",i)
	}
}

func main()  {
	f("direct")

	go f("goroutines")  //可以理解为线程1

	go func(msg string) {
		fmt.Println(msg)
	}("going")    //可以理解为线程2

	fmt.Scanln()    //当程序执行到这行时，程序会停止执行等待用户输入
	fmt.Println("done")
	//那么这两行代码是什么意思呢
	//在整个代码里的作用是控制代码的结束时间，尝试去掉这两行代码，就明白是什么意思了

	//总结22：大家都知道java实现线程需要继承Thread类或者实现runable接口，这样看来go在并发处理上的确有优势
	//PS：好吧，goroutine有专门的名字，叫协程
}
