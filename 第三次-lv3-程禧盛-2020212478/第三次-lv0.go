package main

import (
	"fmt"

)

var (
	myres = make(map[int]int, 20)

)
var channel=make(chan int)
func factorial(n int) {
	var res = 1
    var send int
	for i := 1; i <= n; i++ {
		res *= i
	}
    channel<-send
	myres[n] = res

}

func main() {

	for i := 1; i <= 20; i++ {
		go factorial(i)
		<-channel
	}

	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}


}
