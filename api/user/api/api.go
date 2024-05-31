package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	users "gwf/api/user/mod"
	"gwf/api/user/service"
	"gwf/global"
	"gwf/utils"
)

func Register(c *gin.Context) {
	var user users.Params
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, global.GinBackData(7, "参数不正确", nil))
		return
	}

	global.Redis.Del(user.Email)
	//user.Code
	user.ID, _ = uuid.NewUUID()
	//判断手机号码是否重复
	finder, _ := service.Find("phone", user.Phone)
	if finder.Phone != "" {
		c.JSON(200, global.GinBackData(7, "手机号码已被注册", nil))
		return
	}
	_, err = global.Redis.Get(user.Email).Result()
	if err != nil {
		c.JSON(200, global.GinBackData(7, "邮箱不匹配", nil))
		return
	}
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(200, global.GinBackData(7, "创建失败，请稍后重试", nil))
		return
	}
	var userParams users.Info
	userParams.ID = user.ID
	userParams.Name = user.Name
	userParams.Phone = user.Phone
	userParams.Email = user.Email
	userParams.UserName = user.UserName
	userParams.Sex = user.Sex
	userParams.Password = hashPassword
	userParams.Role = "user"
	err = service.Create(&userParams)
	if err != nil {
		c.JSON(200, global.GinBackData(7, "创建失败，请稍后重试", nil))
		return
	}
	c.JSON(200, global.GinBackData(200, "创建成功", nil))
}

func Login(c *gin.Context) {
	var loginParams users.Login
	err := c.BindJSON(&loginParams)
	if err != nil {
		c.JSON(200, global.GinBackData(7, "参数不正确", nil))
		return
	}
	info, msg := service.Login(loginParams.Email, loginParams.Password)
	if info == nil {
		c.JSON(200, global.GinBackData(7, msg, nil))
		return
	}
	c.JSON(200, global.GinBackData(0, msg, info))
	return
}
