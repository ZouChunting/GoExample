package main

import (
	"fmt"
	"time"
)

//48.time
func main()  {
	p:=fmt.Println
	now:=time.Now()//获取当前的时间
	p(now)

	then:=time.Date(2009,11,17,20,34,58,651387237,time.UTC)
	p(then)

	p(then.Year())//年
	p(then.Month())//月
	p(then.Day())//日
	p(then.Hour())//时
	p(then.Minute())//分
	p(then.Second())//秒
	p(then.Nanosecond())//纳秒
	p(then.Location())//时区
	p(then.Weekday())//一周的周几
	p(then.Before(now))//是否在当前时间之前
	p(then.After(now))//是否在当前时间之后
	p(then.Equal(now))//是否等于当前时间

	diff:=now.Sub(then)  //计算时间then与当前时间的时间差 83220h9m59.387451763s
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))//计算时间then加上时间diff后的日期
	p(then.Add(-diff))//加个-当然就是减的意思了
	//time包中的Add和Sub的用法，Add用于计算某个时间之前和之后的时间点，Sub用于计算两个时间差
	//总结48：没啥水花，各大数据库也有这种功能
}
//2019-05-17 16:44:58.038839 +0800 CST m=+0.000190414
//2009-11-17 20:34:58.651387237 +0000 UTC
//2009
//November
//17
//20
//34
//58
//651387237
//UTC
//Tuesday
//true
//false
//false
//83220h9m59.387451763s
//83220.16649651437
//4.993209989790862e+06
//2.9959259938745177e+08
//299592599387451763
//2019-05-17 08:44:58.038839 +0000 UTC
//2000-05-21 08:24:59.263935474 +0000 UTC
