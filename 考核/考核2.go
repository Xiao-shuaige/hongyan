package main

import (
	"fmt"
	"time"
)

func perfectNumber(a int)  {
	var sum int
	var j int
	var i int
	for j=1;j<a;j++{
		sum=0
		for i=1;i<=j/2;i++{
			if j%i==0{
			sum+=i}
	}
		if j==sum{
			fmt.Printf("%d is a perfect number\n",j)}
	}
}
func main()  {
	var a =123456
	go perfectNumber(a)
	time.Sleep(5 *time.Second)
}
