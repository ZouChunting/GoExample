package main

import (
	"fmt"
	"time"
)

//50.时间格式/解析
func main()  {
	p:=fmt.Println
	t:=time.Now()
	p(t.Format(time.RFC3339))   //2019-05-19T00:35:27+08:00
	//time.RFC3339是一种时间格式，这就是把当前时间按这个格式输出

	t1,_:=time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1) //2012-11-01 22:08:41 +0000 +0000
	//time.Parse 把时间字符串转换为Time，时区是UTC时区。
	//"2012-11-01T22:08:41+00:00"的时间格式是RFC3339

	p(t.Format("3:04PM"))  //12:35AM
	//将当前时间按照"3:04PM"格式输出
	p(t.Format("Mon Jan _2 15:04:05 2006"))  //Sun May 19 00:35:27 2019
	//将当前时间按照"Mon Jan _2 15:04:05 2006"格式输出
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))  //2019-05-19T00:35:27.327958+08:00
	//将当前时间按照"2006-01-02T15:04:05.999999-07:00"格式输出
	form:="3 04 PM"
	t2,_:=time.Parse(form,"8 41 PM")
	p(t2)  //0000-01-01 20:41:00 +0000 UTC

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second())
	//2019-05-19T00:35:27-00:00
	ansic:="Mon Jan _2 15:04:05 2006"
	_,e:=time.Parse(ansic,"8:41PM")
	p(e)  //parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"

}
//关于time的一些格式
//但是具体这些格式有啥讲究，我也不知道
//const (
//	ANSIC       = "Mon Jan _2 15:04:05 2006"
//	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
//	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
//	RFC822      = "02 Jan 06 15:04 MST"
//	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
//	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
//	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
//	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
//	RFC3339     = "2006-01-02T15:04:05Z07:00"
//	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
//	Kitchen     = "3:04PM"
//	// Handy time stamps.
//	Stamp      = "Jan _2 15:04:05"
//	StampMilli = "Jan _2 15:04:05.000"
//	StampMicro = "Jan _2 15:04:05.000000"
//	StampNano  = "Jan _2 15:04:05.000000000"
//)


