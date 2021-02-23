package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"zhihu/models"
)

type CategoriesControllers struct {

}

var CategoriesController CategoriesControllers

func (u *CategoriesControllers)Init()  {
	
}

// 添加分类
func (u *CategoriesControllers)AddCategories(c *gin.Context)  {

	// 声明接收的变量
	var cate models.Categories
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&cate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	id, err := models.CreateCategories(cate, now)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "create failed",
			"err":     err,
		})
		return
	}

	var data struct {
		Id         string `json:"_id"`
		ClassName   string `json:"className"`
		Title    string `json:"title"`
		CreateTime string `json:"createTime"`
	}
	data.Id = strconv.FormatInt(id, 10)
	data.ClassName = cate.ClassName
	data.Title =cate.Title
	data.CreateTime = now
	c.JSON(200, data)

}

func (u *CategoriesControllers)GetCategoriesById(c *gin.Context)  {
	type Param struct {
		Id string `uri:"id" binding:"required"`
	}
	var p Param
	if err := c.ShouldBindUri(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok, cate := models.GetCategoriesById(p.Id)
	if ok {
		c.JSON(200, cate)
	} else {
		c.JSON(403, gin.H{
			"message": "不存在该分类",
		})
	}
}

func (u *CategoriesControllers)GetCategoriesList(c *gin.Context) {
	users := models.GetCategoriesList()
	c.JSON(200, users)
}

func (u *CategoriesControllers)GetCategoriesCount(c *gin.Context) {
	count := models.GetCategoriesCount()
	c.JSON(200, gin.H{
		"categoryCount":count,
	})
}

func (u *CategoriesControllers)UpdateCategoriesById(c *gin.Context)  {
	type UriParam struct {
		Id string `uri:"id" binding:"required"`
	}
	var p1 UriParam
	var p2 models.Categories
	if err := c.ShouldBindUri(&p1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Bind(&p2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, cate := models.UpdateCategoriesById(p1.Id, p2)
	if ok {
		c.JSON(200, cate)
	} else {
		c.JSON(403, gin.H{
			"message": "不存在该用户",
		})
	}
}

func (u *CategoriesControllers)DeleteCategoriesById(c *gin.Context)  {
	type Param struct {
		Id string `uri:"id" binding:"required"`
	}
	var p Param
	if err := c.ShouldBindUri(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok := models.DeleteCategoriesById(p.Id)
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