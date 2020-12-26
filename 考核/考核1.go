package main

import (
	"fmt"
	"strconv"
)
const (
//黑桃
Spade = 0
//红桃
Hearts = 1
//梅花
Club = 2
//方块
Diamond = 3
)

type Poker struct {
Num int
Flower int
}

func (p Poker)PokerSelf()string  {
var buffer string

switch p.Flower {
case Spade:
buffer += "♤"
case Hearts:
buffer += "♡"
case Club:
buffer += "♧"
case Diamond:
buffer += "♢"
}
switch p.Num {
case 13:
buffer += "2"
case 12:
buffer += "A"
case 11:
buffer += "K"
case 10:
buffer += "Q"
case 9:
buffer += "J"
default:
buffer += strconv.Itoa(p.Num+2)
}

return buffer
}

func CreatePokers()(pokers Pokers)  {
for i := 1; i < 14; i++ {
for j := 0; j < 4; j++ {
pokers = append(pokers,Poker{
Num:    i,
Flower: j,
})
}
}
return
}

type Pokers []Poker


func (p Pokers)Print()  {
for _, i2 := range p {
fmt.Print(i2.PrintPoker()," ")
}
fmt.Println()
}
func main()  {
	CreatePokers()
	var p Poker

	fmt.Println(Pokers{})
}
//这题做不来，感觉还是方法和接口还有切片用少了，平时也没怎么用
//思路：洗牌的思路是把切片乱序之后输出
//排序是按照花色用sort排序或者是按照扑克牌大小进行排序，切片排序网上看了看也没学会（逃