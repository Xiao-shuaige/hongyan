package main

import (
	"fmt"
	"strconv"
	"strings"
)

func showMap(m map[string]int)  {
	for k,v:=range m{
		fmt.Println(k,v)
	}
}
func addMap(m map[string]int,key string,value int){
	m[key]=value
}
func add(a *int)  int{
	*a++
	fmt.Println("a加一")
	return *a
}
func multiple(a *int)  int{
	*a *=*a
	fmt.Println("a平方")
	return *a
}
func closepackage()func() int {
	x:=22
	return func() int {
		x/=2
		return x
	}
}
func Atoi(s string)(i int, err error){
	return i,err
}
func Itoa(i int) string{
	var c string
	return   c
}
func ParseFloat(s string, bitSize int) (f float64, err error){
	return f,err
}

func main()  {
	j:=2
	defer fmt.Printf("j的大小为%d\n",j)
	m:=make (map[string]int)
	m["sn"]=2
	m["tes"]=4
	m["dwg"]=1
    showMap(m)
	delete(m,"dwg")
	showMap(m)
	fmt.Println(m)
	key:="fpx"
	value:=100
	addMap(m,key,value)
	j++
	defer fmt.Printf("现在j的大小为%d\n",j)
	fmt.Println(m)
    i:=m["sn"]
	fmt.Println(i)
    var a=2
    fmt.Printf("现在a的大小为%d\n",a)
    b:=add(&a)
    c:=multiple(&a)
    fmt.Println(b,c)
    d:=closepackage()
    fmt.Println(d)
    fmt.Println(d())
	var e int
    e=3
    f:=float32(e)
    fmt.Println(f)
    fmt.Println(strconv.Atoi("abc"))
    fmt.Println(strconv.FormatInt(12,2))
    var s string
    var lm int
    s="rng"
    lm=9
    fmt.Println(strconv.ParseFloat(s,lm))
    var str string
    str="xiaocheng"
    fmt.Println(strings.Index(str,"cheng"))
    fmt.Println(strings.IndexByte(str,'l'))
    fmt.Println(strings.Contains("i love you"," "))
    fmt.Println(strings.Contains("nihao","l"))
    l:="hello world"
    fmt.Println(strings.Count(l,"l"))
    fmt.Println(strings.Count(l," "))
    var f1 string
    var f2 string
    var f3 string
    f1="china"
    f2="italy"
    f3="China"
    fmt.Println(strings.Compare(f1,f2))
    fmt.Println(strings.Compare(f1,f3))
    fmt.Printf("%q\n",strings.Split("nihao,zhongguo",","))
    fmt.Printf("%q\n",strings.Split("nihao,zhongguo"," "))
    fmt.Println(strings.Replace("chengxisheng","e","S",2))
    fmt.Println(strings.Replace("chengxisheng","g","t",-1))
}