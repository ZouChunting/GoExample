package main

import (
	"fmt"
	"time"
)

//33.tickers 滴，滴答器？？？

//ticker会重复做某件事，直到我们停止它
func main()  {
	ticker:=time.NewTicker(500*time.Millisecond)
	go func() {
		for t:=range ticker.C{    //同样ticker也有自己的通道，存入时间节点
			fmt.Println("Tick at",t)
		}
	}()
	time.Sleep(1600*time.Millisecond) //聪明的小机灵鬼应该一眼就看出为什么这样设置时间，保证ticker执行3次
	ticker.Stop() //叫停ticker
	fmt.Println("Ticker stopped")
}
//总结33：这概念很好理解
