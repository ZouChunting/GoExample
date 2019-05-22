package main

import "fmt"
import s "strings"  //这个我知道，python中也有给导入包命名的操作
//44.string
//早知道会讲这个，我之前多此一举搜它干嘛

var p=fmt.Println   //还有这种操作？！我服了

func main()  {
	p("Contains:",s.Contains("test","es"))   //Contains: true
	//返回目标字符串红是否包含子字符串，返回true/false
	p("Count:",s.Count("test","t"))         //Count: 2
	//返回目标字符串中包含的子字符串个数，返回int
	p("HasPrefix:",s.HasPrefix("test","te"))   //HasPrefix: true
	//返回目标字符串是否以子字符串开头，返回true/false
	p("HasSuffix:",s.HasSuffix("test","st"))   //HasSuffix: true
	//返回目标字符串是否以子字符串结尾，返回true/false
	p("Index:",s.Index("test","e"))           //Index: 1
	//返回子字符串在目标字符串中第一次出现的位置，返回int
	p("Join:",s.Join([]string{"a","b"},"-"))      //Join: a-b
	//将目标数组中元素以指定方式连接并返回，返回string
	p("Repeat:",s.Repeat("a",5))              //Repeat: aaaaa
	//将目标字符串按指定个数重复，返回string
	p("Replace:",s.Replace("foo","o","0",-1))    //Replace: f00
	//将目标字符串中的某字符转换为另一字符，-1代表全部转换
	p("Replace:",s.Replace("foo","o","0",1))    //Replace: f0o
	//1代表只转换第一个
	p("Replace:",s.Replace("fooooo","o","0",2))  //Replace: f00ooo
	//2代表转换前2个，返回string
	p("Split:",s.Split("a-b-c-d-e","-"))   //Split: [a b c d e]
	//将目标字符串以指定方式分割，返回数组
	p("ToLower:",s.ToLower("TEST"))     //ToLower: test
	//将目标字符串全部转为小写，返回string
	p("ToUpper:",s.ToUpper("test"))    //ToUpper: TEST
	//将目标字符串全部转为大写，返回string
	p("Len:",len("hello"))    //Len: 5
	//返回目标字符串的长度，返回int
	p("char:","hello"[1])   //char: 101  //返回的是e的ascii码
	//返回目标字符串指定位置元素的ascii码
}
//总结44：不得不说go的这套api是非常完备且易于调用的
