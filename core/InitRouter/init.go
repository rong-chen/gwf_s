package initrouter

import (
	"github.com/gin-gonic/gin"
	email "gwf/api/email/router"
	logFileGroup "gwf/api/logFileGroup/router"
	user "gwf/api/user/router"
)

type RouterInterface interface {
	InitRouter(*gin.RouterGroup)
}

var routerArray = []RouterInterface{
	new(user.Router),
	new(email.Router),
	new(logFileGroup.Router),
}

func LoadRouter(engine *gin.Engine) {
	router := engine.Group("")
	for _, item := range routerArray {
		item.InitRouter(router)
	}
}
