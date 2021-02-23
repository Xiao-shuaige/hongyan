package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"zhihu/models"
)

type loginUser struct {
	User    models.Users
	IsLogin bool
}

type UserControllers struct {
	LoginUser loginUser
}

var UserController UserControllers

func (u *UserControllers)Init() {
	u.LoginUser.IsLogin = false
	u.LoginUser.User = models.Users{
		Id:         0,
		NickName:   "",
		Email:      "",
		Password:   "",
		Role:       "",
		Avatar:     "",
		Status:     0,
		CreateTime: "",
	}
}

// 创建用户
func (u *UserControllers)CreateUsers(c *gin.Context) {

	// 声明接收的变量
	var user models.Users
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	id, err := models.CreateUser(user, now)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "create failed",
			"err":     err,
		})
		return
	}

	var data struct {
		Id         string `json:"_id"`
		Email      string `json:"email"`
		NickName   string `json:"nickName"`
		Role       string `json:"role"`
		Avatar     string `json:"avatar"`
		CreateTime string `json:"createTime"`
		Status     int    `json:"status"`
	}
	data.Id = strconv.FormatInt(id, 10)
	data.Email = user.Email
	data.NickName = user.NickName
	data.Role = user.Role
	data.Avatar = user.Avatar
	data.CreateTime = now
	data.Status = user.Status
	c.JSON(200, data)

}

// 用户登陆
func (u *UserControllers)Login(c *gin.Context) {
	// 声明接收的变量
	var user models.Users
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, newUser := models.CheckUser(user)
	if ok {
		u.LoginUser.IsLogin = true
		u.LoginUser.User = newUser
		var data struct {
			Id         string `json:"_id"`
			Email      string `json:"email"`
			NickName   string `json:"nickName"`
			Role       string `json:"role"`
			Avatar     string `json:"avatar"`
			CreateTime string `json:"createTime"`
			Status     int    `json:"status"`
		}
		data.Id = strconv.FormatInt(newUser.Id, 10)
		data.Email = newUser.Email
		data.NickName = newUser.NickName
		data.Role = newUser.Role
		data.Avatar = newUser.Avatar
		data.CreateTime = newUser.CreateTime
		data.Status = newUser.Status
		c.JSON(200, data)
	} else {
		c.JSON(403, gin.H{
			"message": "登录失败",
		})
	}

}

// 退出登陆
func (u *UserControllers)Logout(c *gin.Context) {
	if u.LoginUser.IsLogin == false {
		c.JSON(403, gin.H{
			"message": "当前未登录",
		})
	} else {
		c.JSON(403, gin.H{
			"message": "退出成功",
		})
	}
}

// 获取登陆状态
func (u *UserControllers)LoginStatus(c *gin.Context) {
	if u.LoginUser.IsLogin == false {
		c.JSON(200, gin.H{
			"isLogin": "false",
		})
	} else {
		c.JSON(403, gin.H{
			"isLogin": "true",
			"id":      u.LoginUser.User.Id,
		})
	}
}

// 获取用户列表
func (u *UserControllers)GetUserList(c *gin.Context) {
	users := models.GetUserList()
	c.JSON(200, users)
}

// 根据id查询用户
func (u *UserControllers)GetUserById(c *gin.Context) {
	type Param struct {
		Id string `uri:"id" binding:"required"`
	}
	var p Param
	if err := c.ShouldBindUri(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok, user := models.GetUserById(p.Id)
	if ok {
		c.JSON(200, user)
	} else {
		c.JSON(403, gin.H{
			"message": "不存在该用户",
		})
	}

}

// 根据id修改用户
func (u *UserControllers)UpdateUserById(c *gin.Context) {
	type UriParam struct {
		Id string `uri:"id" binding:"required"`
	}
	var p1 UriParam
	var p2 models.Users
	if err := c.ShouldBindUri(&p1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Bind(&p2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, user := models.UpdateUserById(p1.Id, p2)
	if ok {
		c.JSON(200, user)
	} else {
		c.JSON(403, gin.H{
			"message": "不存在该用户",
		})
	}
}

func (u *UserControllers)DeleteUserById(c *gin.Context) {
	type Param struct {
		Id string `uri:"id" binding:"required"`
	}
	var p Param
	if err := c.ShouldBindUri(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok := models.DeleteUserById(p.Id)
	if ok {
		c.JSON(200, gin.H{
			"message": "删除成功",
		})
	} else {
		c.JSON(403, gin.H{
			"message": "删除失败",
		})
	}
}

func (u *UserControllers)ChangePasswd(c *gin.Context) {
	type Param struct {
		UserPass    string `form:"userPass" binding:"required"`
		NewPass     string `form:"newPass" binding:"required"`
		ConfirmPass string `form:"confirmPass" binding:"required"`
	}
	var p Param
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if p.NewPass != p.ConfirmPass {
		c.JSON(403,gin.H{
			"message":"密码输入不一致，请确认密码",
		})
		return
	}
	if u.LoginUser.IsLogin == false {
		c.JSON(403,gin.H{
			"message":"尚未登录或登录过期",
		})
		return
	}
	ok := models.ChangePasswd(u.LoginUser.User.Id,p.UserPass,p.NewPass)
	if ok {
		c.JSON(403,gin.H{
			"message":"修改成功",
		})
		return
	}else {
		c.JSON(403,gin.H{
			"message":"修改失败，请检查密码是否正确",
		})
		return
	}
}
