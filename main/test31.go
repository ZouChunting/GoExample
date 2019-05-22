package main

import (
	"fmt"
	"os"
)

//45.字符串格式
type point struct {
	x,y int
}//点结构体

func main()  {
	p:=point{1,2}
	fmt.Printf("%v\n",p)   //{1 2}
	fmt.Printf("%+v\n",p)   //{x:1 y:2}
	fmt.Printf("%#v\n",p)   //main.point{x:1, y:2}
	fmt.Printf("%T\n",p)    //main.point
	fmt.Printf("%t\n",true)   //true
	fmt.Printf("%d\n",123)   //123
	fmt.Printf("%b\n",14)   //1110
	fmt.Printf("%c\n",33)    //!
	fmt.Printf("%x\n",456)    //1c8
	fmt.Printf("%f\n",78.9)   //78.900000
	fmt.Printf("%e\n",1234000000.0)  //1.234000e+09
	fmt.Printf("%E\n",1234000000.0)  //1.234000E+09
	fmt.Printf("%s\n","\"string\"")  //"string"
	fmt.Printf("%q\n","\"string\"")   //"\"string\""
	fmt.Printf("%x\n","hex this")   //6865782074686973
	fmt.Printf("%p\n",&p)    //0xc000094000
	fmt.Printf("|%6d|%6d|\n",12,345)  //|    12|   345|
	fmt.Printf("|%6.2f|%6.2f|\n",1.2,3.45)  //|  1.20|  3.45|
	fmt.Printf("|%-6.2f|%-6.2f|\n",1.2,3.45)   //|1.20  |3.45  |
	fmt.Printf("|%6s|%6s|\n","foo","b")  //|   foo|     b|
	fmt.Printf("|%-6s|%-6s|\n","foo","b")  //|foo   |b     |
	s:=fmt.Sprintf("a %s","string")
	fmt.Println(s)     //a string
	fmt.Fprintf(os.Stderr,"an %s\n","error")  //an error   //但是不知道为啥，我的什么都没打印出来
}

//总结45：我的建议就是，这一节了解一下就行，看一遍，以后要用到哪种再说
//其实跟c还蛮像的
