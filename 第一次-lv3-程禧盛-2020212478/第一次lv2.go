package main

import (
	"fmt"

)

func multiple(a,b int) int {
	return a*b
}
func add(a,b int) int {
	return a+b
}
func minus (a,b int) int {
	return a-b
}
func devide (a,b int) int {
	return a/b
}
func main()  {
  var a int
  var b int
  var c byte
  fmt.Scan(&a)
  fmt.Scan(&b)
  fmt.Scanf("%c",&c)
  //fmt.Printf("%d %d %c",a,b,c)
  /*fmt.Scanf("%d",&a)
  fmt.Scanf("%d",&b)
  fmt.Scanf("%c",&c)*/
  if c=='*' {
  	mul:=multiple(a,b)
  	fmt.Println(mul)
  }
  if c=='+'{
  	di:=add(a,b)
  fmt.Println(di)}
  	if c=='-' {
  		mi:=minus(a,b)
  		fmt.Println(mi)
	}
	if c=='/' {
		de:=devide(a,b)
		fmt.Println(de)
	}

}
