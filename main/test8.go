package main

import (
	"errors"
	"fmt"
)

//21.errors包
//go语言使用error类型来返回函数执行过程中遇到的错误，如果返回的error值为nil，则表示未遇到错误，
//否则error会返回一个字符串，用于说明遇到了什么错误
//nil就是无的意思
//下面这段乍一看有点难理解，不着急，仔细看

//--a--
func f1(arg int)(int,error)  {
	if arg==42{  //遇到错误（这里模拟信号为42即为遇到错误）
		return -1,errors.New("can not work with 42")  //返回错误信号和文本信息
	}else {  //未遇到错误
		return arg+3,nil
	}
}
//--a--

//--b--
type argError struct {
	arg int
	prob string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s",e.arg,e.prob)
}

func f2(arg int)(int,error){
	if arg==42{
		return -1,&argError{arg,"can not work with it"}
	}else {
		return arg+3,nil
	}
}
//--b--

//a,b两段实际上实现了同样的功能，a直接使用了errors包，b定义结构体来自定义error
//重点是这个b如何理解，可以看到f2函数的返回值是int和error，对应于-1,&argError{arg,"can not work with it"}
//要将结构体argError作为error返回，必须实现Error()接口，即上述代码的func (e argError) Error() string方法

func main()  {
	for _,i:=range []int{7,42}{  //用7，42去测试
		if r,e :=f1(i);e!=nil{
			fmt.Println("f1 failed:",e)
		}else {
			fmt.Println("f1 worked:",r)
		}
	}
	for _,i:=range []int{7,42}{
		if r,e:=f2(i);e!=nil{
			fmt.Println("f2 failed:",e)
		}else {
			fmt.Println("f2 worked:",r)
		}
	}
	_,e:=f2(42)
	if ae,ok:=e.(*argError);ok{
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
