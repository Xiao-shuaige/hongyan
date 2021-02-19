package main

import "fmt"

type person struct {
	name string
	age int
	gender string//结构体的三个属性：名字年龄和性别
}
type dove struct{
	person
	doveNum int//内嵌结构体，鸽子有person的所有属性外加鸽的次数
}
type repeater struct {
	person
	repeatNum int//与鸽子同理
}
type lemon struct {
	person
	lemonNum int//同理:柠檬
}
type sweet struct {
	person
	sweetNum int//同理:真香
}

func main()  {
	type dove interface {
		gugugu()//鸽
	}
	type repeater interface {
		repeat()//复读
	}
	type lemon interface {
		jealousy()//酸
	}
	type sweet interface {
		pullBack()//真香
	}//以上是定义出四个接口
	var p person
	func (p*person) dove(){
	}
	func (p*person) repeater(){
	}
	func (p*person) lemon(){
	}
	func (p*person) sweet(){
	}//人类实现以上四个接口

}
