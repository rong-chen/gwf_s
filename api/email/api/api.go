package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gwf/api/email/mod"
	"gwf/api/email/service"
	user "gwf/api/user/service"
	"gwf/global"
	"gwf/utils"
	"time"
)

func SendEmailCode(c *gin.Context) {
	type Params struct {
		Email string `json:email form:email`
	}
	var params Params
	err := c.BindJSON(&params)
	if err != nil {
		c.JSON(200, global.GinBackData(200, "参数错误", nil))
		return
	}
	userinfo, err := user.Find("email", params.Email)
	if userinfo.ID != uuid.Nil {
		c.JSON(200, global.GinBackData(200, "邮箱已被注册", nil))
		return
	}
	var mailer service.Mailer
	mailer.C = "1416307833@qq.com"
	mailer.Account = "1416307833@qq.com"
	mailer.Password = "yetlmmkncinzhdjj"
	strCode := utils.GenerateRandomCode(6)
	mailer.HtmlBody = fmt.Sprintf(mod.HtmlBody, strCode)
	mailer.F = "1416307833@qq.com"
	mailer.T = params.Email
	err = mailer.SendEmail()
	//通过redis记录邮箱和参数，有效时间5分钟
	_, err = global.Redis.Set(params.Email, strCode, 5*time.Minute).Result()
	if err != nil {
		c.JSON(200, global.GinBackData(200, "邮件发送失败，请稍后再试", nil))
		return
	}
	c.JSON(200, global.GinBackData(200, "发送成功", nil))
	return
}
