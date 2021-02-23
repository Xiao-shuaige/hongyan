package main

import (
	"zhihu/controllers"
	"zhihu/models"
	"zhihu/routers"
)


func main() {

	models.InitDB()
	initControllers()

	routers.InitRouters()
}

func initControllers()  {
	// 初始化用户控制器
	controllers.UserController.Init()
	// 初始化分类控制器
	controllers.CategoriesController.Init()
}