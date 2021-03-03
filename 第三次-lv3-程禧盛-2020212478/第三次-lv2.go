package main

import (
	"fmt"
	"os"
)
func main()  {
	os.Create("proverb.txt")
	fi,err:=os.OpenFile("c:/hello.world/proverb.txt",os.O_WRONLY,6)
	if err!= nil{
		fmt.Println("文件创建失败")
		return
	}
	defer fi.Close()
	fi.WriteString("don't communicate by sharing memory share memory by communicating")
	fp,err:=os.Open("c:/hello.world/proverb.txt")
	if err!=nil{
		fmt.Println("文件打开失败")
		return
	}
	b:=make([]byte,1000)
	fp.Read(b)
	for i:=0;i<len(b);i++{
		if b[i]!=0{
			fmt.Printf("%c",b[i])
		}
	}
	defer fp.Close()
}