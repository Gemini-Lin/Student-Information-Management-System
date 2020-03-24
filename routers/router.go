package routers

import (
	"StudentSystem/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	// login 采用中间件设置
	loginGroup := r.Group("login")
	{
		// 学生信息
		// 添加
		//本来是想要采用post，get，put,delete的restful api风格，但是发现from表单提交时只能返回post,get方法
		loginGroup.POST("/Student", controller.CreateStudent)
		// 查看所有的学生信息
		loginGroup.GET("/Student", controller.GetStudentList)
		// 修改某一个学生信息
		loginGroup.GET("/Student/modify", controller.UpdateAStudent)
		// 删除某一个学生信息
		loginGroup.GET("/Student/delete", controller.DeleteAStudent)
	}
	return r
}