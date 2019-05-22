package main

import "fmt"

//30.关闭通道
//其实关闭通道本身并没有什么特别要讲的，但是作为一个学艺不精的，不禁有个不成问题的问题
//go本身也是有gc的，那通道回收在gc射程内吗？为什么要有手动关闭？
func main()  {
	jobs:=make(chan int,5)
	done:=make(chan bool)

	go func(){
		for{
			j,more:=<-jobs  //这里让人有点恍惚，通道里到底存个啥
			//关于这个，可以尝试一下下面几行代码
			//ch:=make(chan string,5)
			//	ch<-"12345"
			//	ch<-"asdfg"
			//	i,j:=<-ch
			//	fmt.Println("i",i)    "12345"
			//	fmt.Println("j",j)    true
			//	m:=<-ch
			//	fmt.Println("m",m)    "asdfg"
			//	n:=<-ch
			//	fmt.Println("n",n)    报错
			//简单讲一下个人解释，缓存通道中的存储的每个元素，都有一个元素标签
			//每次读值i,j:=<-ch，i是元素的值，j是元素标签，为true的话代表存在有效值，不存在有效值的话应该是nil，这个我还没验证
			//而m:=<-ch，n:=<-ch这种写法，导致第二句报错的原因，我想是值被读出后，元素标签自然消失

			if more{
				fmt.Println("more",more)
				fmt.Println("received job",j)
			}else {
				fmt.Println("received all jobs")
				done<-true
				return
			}
		}
	}()
	//有没有人想过，为什么这里用for来写，可不可以用select-default呢，下面我写了一段
	//报错fatal error: all goroutines are asleep - deadlock!
	//我水平不够，想不到有什么方法可以用select-default写出来
	//但我想这段代码的要点是，我们必须使用到more来确定通道中是否存在有效值，才能读值打印
	//否则会报错。
	//感觉这个解释不是很到位，看到后面会有更好的解释吧


	for j:=1;j<=3;j++{
		jobs<-j
		fmt.Println("sent job",j)
	}//简单的向通道写值，1，2，3
	close(jobs)
	fmt.Println("sent all jobs")

	<-done   //必须等待done值为true，才能退出
	//可以试着把这行去掉，多执行几次代码，能够发现协程能否执行完毕完全是玄学
	//可以先这么理解，main作为一个主线程，它不会等待其他线程执行完毕再退出，而是主线程执行完毕
	//其他线程也强制结束，因此加上<-done，胁迫主线程必须等待done值，也就是完成信号，才能退出
	//并且，我们可以多想一步，如果没有<-done，哪些输出语句是必定可以执行的，而哪些是玄学
	//显而易见，主线程里的所有sent都可以被执行打印，而go协程里的received能否执行完毕看脸。
}
//总结30：关于这部分的附加代码

//jobs:=make(chan int,5)
//done:=make(chan bool)
//
//go func() {
//	select {
//	case j:=<-jobs:
//		fmt.Println("received",j)
//	default:
//		fmt.Println("received all")
//		done<-true
//	}
//}()
//
//for i:=1;i<=3;i++{
//jobs<-i
//fmt.Println("sent",i)
//}
//close(jobs)
//fmt.Println("sent all jobs")
//
//<-done

//ch:=make(chan string,5)
//ch<-"12345"
//ch<-"asdfg"
//done:=make(chan bool)
//go func() {
//	for {
//		_,v:=<-ch
//		if v != true {
//			fmt.Println("there is no values")
//			done<-true
//		}else {
//			fmt.Println("lalalla")
//			done<-true
//		}
//	}
//}()
//<-done