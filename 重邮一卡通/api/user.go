package api

import (
	"github.com/gin-gonic/gin"
	"main.go/model"
	"main.go/service"
	"net/http"
)

// 注册接口
func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.String(http.StatusInternalServerError, "用户名或密码为空")
		return
	}
	flag := service.CheckUser(username)
	if flag == true {
		c.String(500, "用户名已存在")
		return
	}
	err := service.CreateUser(model.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		c.String(500, "创建新用户失败，请重试")
		return
	}
	c.String(200, "注册成功")
}

// 登录接口
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.String(http.StatusInternalServerError, "用户名或密码为空")
		return
	}
	flag := service.CheckUser(username)
	if flag == false {
		c.String(500, "用户名不存在")
		return
	}
	flag = service.CheckPassword(username, password)
	if flag == false {
		c.String(500, "密码错误")
		return
	}
	c.SetCookie("user", username, 3600, "/", "127.0.0.1", false, true)

	c.JSON(200, gin.H{
		"message": "登录成功",
	})
}
