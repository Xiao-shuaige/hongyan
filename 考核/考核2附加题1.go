package main

import (
	"fmt"
	"time"
)

var flag int
func primeNumber(a int) {
	var i int

	var j int
	for j = 2; j <= a; j++ {
		flag = 1
		for i = 2; i < j; i++ {
			if j%i == 0 {
				flag = 0
			}
		}
		if flag == 1 {
			fmt.Printf("%d is a primenumber\n", j)
		}
	}
}

func main()  {
	var a=123456
	go primeNumber(a)
	time.Sleep(40 *time.Second)
}
