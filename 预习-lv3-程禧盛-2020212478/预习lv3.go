package main


import "fmt"

func main()  {
	var a[5] int
	var max int
	max=0
	for i:=0;i<5;i++{
		fmt.Scan(&a[i])
		if a[i]>max{
			max=a[i]
		}
	}
	fmt.Println(max)
}
