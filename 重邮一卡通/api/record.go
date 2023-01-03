package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/model"
	"main.go/service"
	"strconv"
	"time"
)

// 备注接口
func attach(c *gin.Context) {
	username, err := middleware.CheckLogin(c)
	if err != nil {
		c.String(401, "未登录！无法访问！")
		return
	}
	uid := service.GetUId(username)
	strRid := c.PostForm("rid")
	context := c.PostForm("context")
	rid, _ := strconv.Atoi(strRid)
	r, err := service.GetRecordByRid(rid)
	if err != nil {
		c.String(500, "获取Record失败")
		return
	}
	if uid != r.Uid {
		c.String(402, "您没有权限查看他人交易记录")
		return
	}
	err = service.AttachContext(rid, uid, context)
	if err != nil {
		c.String(500, "修改备注失败")
		return
	}
	c.String(200, "修改备注成功")
}

func FindRecord(c *gin.Context) {
	username, err := middleware.CheckLogin(c)
	if err != nil {
		c.String(401, "未登录！无法访问！")
		return
	}
	context := c.PostForm("context")
	uid := service.GetUId(username)
	r, err := service.GetRecordByContext(uid, context)
	if err != nil {
		c.String(500, "未能查询到相应交易记录")
		return
	}
	c.String(200, "交易记录如下")
	for i := 0; i < 5; i++ {
		c.JSON(200, model.Record{
			Rid:       r[i].Rid,
			Uid:       r[i].Uid,
			CreatedAt: time.Time{},
			Context:   r[i].Context,
			Money:     r[i].Money,
			Object:    r[i].Object,
		})
	}
}
