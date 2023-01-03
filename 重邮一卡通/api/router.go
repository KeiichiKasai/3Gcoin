package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS()) //解决跨域
	u := r.Group("/user")    //用户组
	{
		u.POST("/register", register)
		u.POST("/login", login)
	}
	r.POST("/charge", charge)
	r.POST("/transfer", transfer)
	r.POST("/attach", attach)
	r.POST("/find", FindRecord)
	r.Run()
}
