package main

import (
	"fmt"
)

func main()  {
	var x float64
	var y float64
	     for y=1.5;y>=-1.5;y-=0.1{
	     	for x=-1.5;x<=1.5;x+=0.1{
				if (x*x+y*y-1)*(x*x+y*y-1)*(x*x+y*y-1)-x*x*y*y*y<=0.0{
					fmt.Printf("*")
				}else {
					fmt.Printf(" ")
				}}
			fmt.Println( )}


	}

