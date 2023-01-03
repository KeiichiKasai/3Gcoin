package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// CheckLogin 检测是否存在cookie
func CheckLogin(c *gin.Context) (username string, err error) {
	username, err = c.Cookie("user")
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	return username, err
}
