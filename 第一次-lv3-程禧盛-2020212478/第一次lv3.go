package main

import (
	"fmt"
)
import "math"

func max(a[100]float64) float64 {
	var n int
	fmt.Scan(&n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	var max float64
	max=0
	for i:=0;i<n;i++{
		if math.Abs(a[i])>max {
			max=math.Abs(a[i])
		}
	}
	return max
}
func main()  {
	var a [100] float64
	b:=max(a)
	fmt.Println(b)

}
