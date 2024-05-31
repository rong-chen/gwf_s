package router

import (
	"github.com/gin-gonic/gin"
	"gwf/api/logFileGroup/api"
)

type Router struct {
}

func (Router) InitRouter(rg *gin.RouterGroup) {
	router := rg.Group("/file")
	{
		router.POST("/add", api.Create)
		router.GET("/list", api.GetList)
	}
}
