package main

import (
	"fmt"
)

func main() {
	over := make(chan bool,10)//使用带缓存的channel
	for i := 0; i < 10; i++ {
		 func() {
			fmt.Println(i) //用函数不用协程
		}()
		if i == 9 {
         over<-true
         close(over)//关闭管道
		}
	}
	<-over
	fmt.Println("over!!!")

}