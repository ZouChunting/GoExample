package main

import (
	"bytes"
	"fmt"
	"regexp"
)

//46.正则表达式
//？喵喵喵，其大名如雷贯耳，我这种萌新就先跳过了
//不过还是把代码贴在这

func main()  {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
//true
//true
//peach
//[0 5]
//[peach ea]
//[0 5 1 3]
//[peach punch pinch]
//[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
//[peach punch]
//true
//p([a-z]+)ch
//a <fruit>
//a PEACH