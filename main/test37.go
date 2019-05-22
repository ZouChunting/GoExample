package main

import (
	"fmt"
	"math/rand"
	"time"
)

//51.随机数
func main()  {
	fmt.Print(rand.Intn(100),",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5,",")
	fmt.Print((rand.Float64()*5)+5)
	fmt.Println()

	s1:=rand.NewSource(time.Now().UnixNano())
	r1:=rand.New(s1)
	fmt.Print(r1.Intn(100),",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
	//前面水了那么久，这里终于有个重点
	//多运行几次程序可以发现
	//fmt.Print(rand.Intn(100),",")
	//fmt.Print(rand.Intn(100))
	//这两句产生的随机数，在产生一次后就固定了，比如我这里产生的是81,87，不管重新运行多少次，它还是81,87，
	//甚至关掉项目重来，它还是81,87，随机数在随机一次后就不随机了，因此它又被叫做伪随机数
	//而fmt.Print(r1.Intn(100),",")
	//	fmt.Print(r1.Intn(100))
	//这两行的随机数就一直是随机的，每次运行都不一样

	s2:=rand.NewSource(42)
	r2:=rand.New(s2)
	fmt.Print(r2.Intn(100),",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3:=rand.NewSource(42)
	r3:=rand.New(s3)
	fmt.Print(r3.Intn(100),",")
	fmt.Print(r3.Intn(100))
	//种子一样，产生的随机数也一样
	//所以我们上面为什么这样写呢？s1:=rand.NewSource(time.Now().UnixNano())
	//就是以时间为种子，以确保随机数不一样



}
