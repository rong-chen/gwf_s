package user

import (
	"gwf/api/user/api"

	"github.com/gin-gonic/gin"
)

type Router struct{}

func (Router) InitRouter(r *gin.RouterGroup) {
	router := r.Group("/user")
	{
		router.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, &gin.H{
				"code": 0,
				"msg":  "success",
			})
		})
		router.POST("/create", api.Register)
		router.POST("/login", api.Login)
	}
}
