package router

import (
	"github.com/gin-gonic/gin"
	"gwf/api/email/api"
)

type Router struct{}

func (Router) InitRouter(r *gin.RouterGroup) {
	router := r.Group("/email")
	{
		router.POST("/sendCode", api.SendEmailCode)
	}
}
