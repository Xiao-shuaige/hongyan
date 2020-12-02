package main

import "fmt"

type common struct {
	name string
	vip bool
	number int
	signature string
}
type author struct{
	a common
	coin int
	like int
}
type watcher struct{
	b common
    love bool
    vote bool
}
func main()  {
	p:=author{
		a: common{
			"xiaocheng",
			true,
			10000,
			"happy is the most important thing",
		},
		coin: 2020,
		like: 10086,
	}
	fmt.Println(p)
	q:=watcher{
		b: common{
			"tianbeibei",
			false,
			100,
			"the most important thing is eating!",
		},
		love: false,
		vote: false,
	}
	fmt.Println(q)
}
