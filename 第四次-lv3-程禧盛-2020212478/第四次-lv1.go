package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	valueByKey =make(map[string]int)
	valueByKeyGuard sync.Mutex
)
func main()  {
	m:=make(map[string]int)//用Make声明初始化一个map类型的变量
	m["xiaocheng"]=1
	m["xiaohu"]=2
	m["xiaotian"]=3    //进行三次赋值操作
	for key,value:=range m{
		fmt.Println(key,value) //遍历索引和值
	}
	fmt.Println("the value of m[\"xiaocheng\"] is ",m["xiaocheng"])//查询一个在Map中已有的值
	value,ok:=m["c"]
	fmt.Println(value,ok)//查询一个map中没有的值，当ok值为false的时候即map不存在这个索引
	i:=1
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)//通过这次实验我们能得出的结论：1.defer紧跟着的语句是在其他语句之后再执行的
	//2.defer存储数据的栈和数据结构中的栈类似，先进后出
	delete(m,"xiaocheng")//使用delete函数删除map中的一个索引
	fmt.Println(m)
    callMap(m,"xiaowang",5)//调用callMap函数
    fmt.Println(m)
	f:= func() {
		fmt.Println("加油打工人")
	}
	f()//使用匿名函数打印出”加油打工人“
	func (a,b int){
	fmt.Println("the result of a+b is ",a+b)
	}(3,4)//再试试这种直接引用的匿名函数
	visit([] int{1,2,3,4}, func(v int){
	   fmt.Println("打印v的值",v)
	})//将匿名函数进行封装，其中函数参数也是一个函数
	fmt.Println("the first call the fuction", multiple())
	fmt.Println("the second call the fuction", multiple())//两次打印同一个函数
	f1:=quadratic()
	fmt.Println("the first call the closure",f1())
	fmt.Println("the second call the closure",f1())
	fmt.Println("the third call the closure",f1())//三次打印同一个闭包
	str:="hello world"
	foo:= func() {
		str="hello google"//匿名函数中并没有定义str,str的定义在匿名函数之前，此时，str就被引用到了匿名函数中形成了闭包
	}
	foo()
	fmt.Println(str)
	str="hello wechat"
	fmt.Println(str)//但是在匿名函数外部仍然可以对str变量进行修改
	generator:=playerNature("xiaocheng")//创建一个玩家生成器
	name,hp:=generator()//返回玩家的名字和血量
	fmt.Println("id：",name,"生命值为",hp)
	a:=fileSize("filename")
	fmt.Println("the return of fileSize fiction is ",a)
	funcB()
    fmt.Println(strconv.ParseInt("0xFF",0,0))//这是个字符串转整数类型的内置函数，参数分别是string,base:想转换成的进制数，
    //bitsize：结果必须能无溢出赋值的整数类型
    fmt.Println(strconv.FormatInt(12,10))//这是个整数转字符串的内置函数，里面有两个参数，第一个是int，第二个是base：像转换成的进制数
    v:="4.12345678"
    if s,err:=strconv.ParseFloat(v,32); err==nil{//字符串转浮点数
         fmt.Printf("%T,%v\n",s,s)//如果符合语法规则，函数会返回一个最接近的浮点数
	}
	//strconv系的东西就先写这么多
	st:="hello world"
	fmt.Println("第一次出现e的位置是：",strings.Index(st,"e"))//查找这个字符串在字符串中第一次出现的位置，不存在则返回-1
	fmt.Println("如果是-1的话就说明没出现过这个字符串",strings.Index(st,"ch"))
	//其他的系就不操作了，是在是太多了
}
func callMap(m map[string]int,key string,value int)  {//声明一个函数，这个函数能使map中多一对索引与值
	m[key]=value
}
func multiple() int  {
	a:=2
	a++
	return a*a
}
func quadratic() func() int  {
	x:=0
	return func() int {
		x++
		return x*x
	}
}
func visit(slice []int,f func(int)){
	for _,v:=range slice {
		 f(v)
	}
}
func playerNature(name string) func() (string,int){
	hp:=100 //血量一直为100
	return func() (string, int) {
		return name,hp //将变量引用到闭包中
	}//返回创建的闭包
}
func fileSize(filename string) int64  {
	f,err:=os.Open(filename)
	if err!=nil {
		return 0
	}
	info,err:=f.Stat()
	if err!=nil {
		f.Close()
		return 0
	}
	size:=info.Size()
	f.Close()
	return size
}

func funcB() {
	defer func(){
		//捕获panic，并恢复程序使其继续运行
		if err := recover(); err != nil {
			fmt.Println("recover in funcB")
		}
	}()

	panic("panic in B")  //主动抛出异常
}

