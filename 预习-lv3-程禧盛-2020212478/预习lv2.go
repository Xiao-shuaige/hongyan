package main

import (
	"fmt"
)

func main()  {
	var a [3]string
	a[0]="2020212478"
	a[1]="2020212480"
	a[2]="2020212497"
	var b[3]string
	b[0]="115319"
	b[1]="123456"
	b[2]="030416"
	var c[3]string
    for i:=0;i<3;i++{
    	fmt.Println("请输入账号")
    	fmt.Scan(&a[i])
    	fmt.Println("请输入密码")
    	fmt.Scan(&c[i])
    	if b[i]==c[i]{
    	fmt.Println("密码正确")}else
    	{
    	fmt.Println("密码错误")}
	}

}
