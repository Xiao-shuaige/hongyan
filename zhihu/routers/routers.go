package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhihu/controllers"
)

func InitRouters()  {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	// 创建用户
	r.POST("/users",controllers.UserController.CreateUsers)
	// 用户登陆
	r.POST("/login",controllers.UserController.Login)
	// 退出登陆
	r.POST("/logout",controllers.UserController.Logout)
	// 获取用户登录状态
	r.GET("/login/status",controllers.UserController.LoginStatus)
	// 获取用户列表
	r.GET("/users",controllers.UserController.GetUserList)
	// 根据id查询用户
	r.GET("/users/id/:id",controllers.UserController.GetUserById)
	// 根据id修改用户
	r.PUT("/users/id/:id",controllers.UserController.UpdateUserById)
	// 根据id删除用户
	r.DELETE("/users/id/:id",controllers.UserController.DeleteUserById)
	// 修改用户密码
	r.PUT("/users/password",controllers.UserController.ChangePasswd)


	// 添加分类
	r.POST("/categories",controllers.CategoriesController.AddCategories)
	// 根据id查询分类
	r.GET("/categories/id/:id",controllers.CategoriesController.GetCategoriesById)
	// 查询分类列表
	r.GET("/categories",controllers.CategoriesController.GetCategoriesList)
	// 查询分类数量
	r.GET("/categories/count",controllers.CategoriesController.GetCategoriesCount)
	// 修改分类
	r.PUT("/categories/id/:id",controllers.CategoriesController.UpdateCategoriesById)
	// 删除分类
	r.DELETE("/categories/id/:id",controllers.CategoriesController.DeleteCategoriesById)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}