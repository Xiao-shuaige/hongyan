package main

import (
	"fmt"
	"time"
)

func main() {
	over := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			over <- true
		}()
		time.Sleep(time.Millisecond)//第一个问题，死锁，因为无缓冲的channel在Main函数里面传递信息，会造成死锁，把他放在协程里
	}//第二个问题，跑太快了，每个循环里面都让main函数调用休眠一段时间，避免协程打不出前几个值
    <-over
	fmt.Println("over!!!")
}