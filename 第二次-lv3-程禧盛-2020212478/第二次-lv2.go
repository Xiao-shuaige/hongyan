package main

import "fmt"

func receiver(v interface{}) {//interface{}空接口，v可以是任意类型，但其实它是 interface{} 类型。
	switch v.(type){       //v.(type)只能在switch中使用，意义：将v的可能的基类型当作表达式使用
	case string:
		fmt.Println("这个是string")
		break
	case int:
		fmt.Println("这个是int")
		break
	case byte:
		fmt.Println("这个是byte")
		break
		//case float:
		fmt.Println("这个是float")
		break
	case bool:
		fmt.Println("这个是bool")
		break
	default:
		fmt.Println("未知类型")
		break
	}}
func main()  {
	a:=32
	b:=28.8
	c:='c'
	d:="nihao"
	e:=true
	receiver(a)
	receiver(b)
	receiver(c)
	receiver(d)
	receiver(e)//以上是几种测试
}
