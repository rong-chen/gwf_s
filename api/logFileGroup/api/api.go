package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gwf/api/logFileGroup/mod"
	"gwf/api/logFileGroup/service"
	"gwf/global"
)

func Create(c *gin.Context) {
	var data mod.LogFileGroup
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(200, global.GinBackData(7, "创建失败,请检查参数", nil))
		return
	}
	uid, _ := uuid.NewUUID()
	data.ID = uid
	err = service.CreateRow(&data)
	if err != nil {
		c.JSON(200, global.GinBackData(7, "创建失败,请检查参数", nil))
		return
	}
	c.JSON(200, global.GinBackData(0, "", data))
}

func GetList(c *gin.Context) {
	id, ok := c.GetQuery("dirID")
	cid, cok := c.GetQuery("createUserId")

	if !cok {
		c.JSON(200, global.GinBackData(7, "参数错误", nil))
		return
	}
	if !ok {
		c.JSON(200, global.GinBackData(7, "参数错误", nil))
		return
	}
	data := service.SearchDir(id, cid)
	c.JSON(200, global.GinBackData(0, "查询成功", data))
}
