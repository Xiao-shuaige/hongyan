package main

import (
	"fmt"
	"time"
)

func main()  {
	go printTheNumber()
	go printTheNumber()
	time.Sleep(time.Millisecond)
}
func printTheNumber()  {
	for i:=1;i<101;i++{
		fmt.Println(i)
	}

}
