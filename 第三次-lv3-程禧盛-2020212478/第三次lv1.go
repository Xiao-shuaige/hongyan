package main

import "fmt"

func add()  {
	for i:=0;i<100;i++{
		fmt.Println(i)
	}
}
/*func repeat()  {
	for i:=0;i<100;i++{
		fmt.Println(i)
	}
}*/
func main()  {
	go add()
	add()
}
