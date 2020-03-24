package controller

import (
	"StudentSystem/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*设计逻辑
 url     --> controller  --> logic   -->    model
获取请求  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

//主页面
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//增加数据
func CreateStudent(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var Student models.Student
	c.ShouldBind(&Student) //获取Json数据
	// 2. 存入数据库
	err:=models.CreateAStudent(&Student)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		//返回状态码及存入数据的json文件
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg": "success",
			"data": Student,
		})
	}
}

//查询数据
func GetStudentList(c *gin.Context) {
	// 查询Student这个表里的所有数据
	StudentList, err := models.GetAllStudent()
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, StudentList)
	}
}

//更新数据
func UpdateAStudent(c *gin.Context) {
	UID, ok := c.GetQuery("UID")//根据学号查询
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "输入的学号有误"})
		return
	}
	Student, err := models.GetAStudent(UID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.ShouldBind(&Student)
	if err = models.UpdateAStudent(Student); err!= nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, Student)
	}
}

//删除数据
func DeleteAStudent(c *gin.Context) {
	UID, ok := c.GetQuery("UID")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "输入的学号有误"})
		return
	}
	if err := models.DeleteAStudent(UID);err!=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{UID:"deleted"})
	}
}