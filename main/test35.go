package main

import (
	"fmt"
	"time"
)

//49.Epoch  纪元？？？
func main()  {
	now:=time.Now() //获取当前时间，返回Time类型
	secs:=now.Unix()  //返回时间戳，自从1970年1月1号到现在
	nanos:=now.UnixNano()//返回时间戳，包含纳秒
	fmt.Println(now)  //2019-05-17 16:51:10.288762 +0800 CST m=+0.000206360

	millis:=nanos/1000000   //微秒
	fmt.Println(secs)//1558083070
	fmt.Println(millis)//1558083070288
	fmt.Println(nanos)//1558083070288762000
	fmt.Println(time.Unix(secs,0))//2019-05-17 16:51:10 +0800 CST
	fmt.Println(time.Unix(0,nanos))//2019-05-17 16:51:10.288762 +0800 CST
	fmt.Println(time.Unix(0,0))//1970-01-01 08:00:00 +0800 CST
	//unix中的这两个参数sec,nsec分别代表秒和纳秒
	//有没有人好奇，什么是时间戳
	//unix时间戳是从1970年1月1日（UTC/GMT的午夜）开始所经过的秒数，不考虑闰秒。
	//关于time包的Unix，可以以一个有效时间调用它，或者直接调用，但参数必不可少，否则会报错
	//至于time.Unix(secs,0)什么意思，结合上面3行代码，其实很好理解
	//就是1970-01-01 08:00:00 +0800 CST经过secs后的时间

	//另外，time.Duration 类型代表两个时间点之间经过的纳秒数，可表示的最长时间段约为290年。
}
//参考链接：https://www.jianshu.com/p/9d5636d34f17
//https://studygolang.com/articles/240
//总结49：
//获取一个时间的方法：
//func Now() Time {} // 当前本地时间
//func Unix(sec int64, nsec int64) Time {} // 根据时间戳返回本地时间
//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time {} // 返回指定时间

//so?这还是在讲时间啊，和纪元有什么关系





