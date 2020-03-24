package main

import (
	"StudentSystem/database"
	"StudentSystem/models"
	"StudentSystem/routers"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE StudentSystem;
	// 连接数据库
	err := database.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer database.Close()  // 程序退出关闭数据库连接
	// 模型绑定
	database.DB.AutoMigrate(&models.Student{})
	// 注册路由
	r := routers.SetupRouter()
	r.Run()//默认在127.0.0.1:8080端口运行
}
