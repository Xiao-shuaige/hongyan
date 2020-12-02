package main

import "fmt"

type construction struct {
	 a int
	 b float64
	 c byte
	 d string
}

func main()  {
	var p construction
    p.a=20
    p.b=18.8
    p.c='^'
    p.d="beibei"
    fmt.Println(p.a,p.b,p.c,p.d)
    fmt.Println(p)
    b:=construction{
    	a:20,
    	b:19.9,
    	c:'*',
    	d:"zhongguo",
	}
	fmt.Println(b)
}
