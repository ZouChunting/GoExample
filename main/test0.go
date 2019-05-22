package main

import (
	"fmt"
	"strings"
)

func main()  {
	//for _,i:=range []int{7,10}{
	//	fmt.Println(i)
	//}
	//ch:=make(chan string,5)
	//go func() {
	//	ch<-"aaa"
	//}()
	//ch<-"12345"
	//ch<-"asdfg"
	//i,j:=<-ch
	//fmt.Println("i",i)
	//fmt.Println("j",j)
	//m:=<-ch
	//fmt.Println("m",m)
	//n:=<-ch
	//fmt.Println("n",n)

	//done:=make(chan bool)
	//	//go func() {
	//	//	for {
	//	//		_,v:=<-ch
	//	//		if v != true {
	//	//			fmt.Println("there is no values")
	//	//			done<-true
	//	//		}else {
	//	//			fmt.Println("lalalla")
	//	//			done<-true
	//	//		}
	//	//	}
	//	//}()
	//	//
	//	//<-done





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
	//	jobs<-i
	//	fmt.Println("sent",i)
	//}
	//close(jobs)
	//fmt.Println("sent all jobs")
	//
	//<-done

	//var mutex = &sync.Mutex{}   //互斥量，如此声明可以加锁或者解锁
	//var i int=0
	//var j int=0
	//
	//go func() {
	//	mutex.Lock()
	//	i+=1
	//	fmt.Println("i1:",i)
	//	//mutex.Unlock()
	//}()
	//
	//go func() {
	//	mutex.Lock()
	//	i+=1
	//	fmt.Println("i2",i)
	//	fmt.Println("ytrew")
	//	j+=1
	//}()
	//
	//time.Sleep(time.Second)
	//fmt.Println("i:",i)
	//fmt.Println("j:",j)

	//ch:=make(chan int)
	//go func() {
	//	ch<-1
	//}()
	//ch<-1

	//fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//fmt.Println(4)

	s:="hello,world!!!!"
	a:=strings.Count(s,"!")
	fmt.Println(a)  //4
	b:=strings.Contains(s,"!!")
	fmt.Println(b)   //true
	c:=strings.ContainsAny(s,"abcde")
	fmt.Println(c)   //true
	e:=strings.Index(s,"l")
	fmt.Println(e)    //2
	f:=strings.LastIndex(s,"l")
	fmt.Println(f)   //9
	g:=strings.HasPrefix(s,"hello")
	fmt.Println(g)   //true
	h:=strings.HasSuffix(s,"!")
	fmt.Println(h)   //true
	i:=strings.ToUpper(s)
	fmt.Println(i)   //HELLO,WORLD!!!!

}
