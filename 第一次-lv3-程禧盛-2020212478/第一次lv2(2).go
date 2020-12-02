package main

import (
	"fmt"
)

func multiple (a, b int) int  {
	return a*b
}
func add(a, b int ) int  {
	return a+b
}
func minus(a,b int) int {
	return a-b
}
func devide(a,b int) int {
	return a/b
}
func main()  {
	var a int
	var b int
	var c byte
	for i:=1;i<100;i++{
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scanf("%c",&c)
	switch c {
	case '*':
		mul:=multiple(a,b)
		fmt.Println(mul)
		break
	case '+':
		ad:=add(a,b)
		fmt.Println(ad)
		break
	case '-':
		mi:=minus(a,b)
		fmt.Println(mi)
		break
	case '/':
		de:=devide(a,b)
		fmt.Println(de)
		break
	}}

}