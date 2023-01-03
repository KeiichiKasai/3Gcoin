package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/model"
	"main.go/service"
	"strconv"
)

// 充值接口
func charge(c *gin.Context) {
	username, err := middleware.CheckLogin(c)
	if err != nil {
		c.String(401, "未登录！ 无法访问！")
		return
	}
	amount := c.PostForm("amount")
	money, _ := strconv.Atoi(amount)
	err = service.AddMoney(username, money)
	if err != nil {
		c.String(500, "充值失败！")
		return
	}
	c.String(200, "充值成功！")
	uid := service.GetUId(username)
	err = service.CreateRecord(model.Record{
		Uid:     uid,
		Context: "为自己充值",
		Money:   money,
		Object:  username,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "创建交易记录失败",
		})
		return
	}

}

// 转账接口
func transfer(c *gin.Context) {
	username, err := middleware.CheckLogin(c)
	if err != nil {
		c.String(401, "未登录！无法访问！")
		return
	}
	target := c.PostForm("target")
	amount := c.PostForm("amount")
	money, _ := strconv.Atoi(amount)
	flag, err := service.MinusMoney(username, money)
	if flag {
		c.String(500, "余额不足，无法给他人转账")
		return
	}
	if err != nil {
		c.String(500, "转账过程出现问题，请重试！")
		return
	}
	err = service.AddMoney(target, money)
	if err != nil {
		c.String(500, "转账过程出现问题，请重试！")
		return
	}
	uid := service.GetUId(username)
	err = service.CreateRecord(model.Record{
		Uid:     uid,
		Context: "给" + target + "转账",
		Money:   money,
		Object:  target,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "创建交易记录失败",
		})
		return
	}
	c.String(200, "转账成功")
}
