package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//36.原子计数器
func main()  {
	var ops uint64
	for i:=0;i<50;i++{
		go func() {
			for{
				atomic.AddUint64(&ops,1)  //递增1的操作
				time.Sleep(time.Millisecond)
			}
		}()
	}//开启了50个协程，每个协程大概执行1ms

	time.Sleep(time.Second)  //可以修改一下时间，这里限制上面50个协程的执行时长
	opsFinal:=atomic.LoadUint64(&ops)
	fmt.Println("ops",opsFinal)
	//程序结果表明了执行操作次数
}
//总结36：这段代码不是很难理解，但是注意指针的运用，对于指针，我已经杯弓蛇影了。注意，add和load的操作
