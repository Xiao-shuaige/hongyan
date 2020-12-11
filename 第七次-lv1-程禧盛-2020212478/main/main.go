package main

import (
	"awesomeProject/cmd"
	"awesomeProject/dao"
)

func main() {
	dao.MysqlInit()
	cmd.Entrance()
}