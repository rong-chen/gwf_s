package main

import (
	initmysql "gwf/core/InitMySQL"
	initrouter "gwf/core/InitRouter"
	viper "gwf/core/InitViper"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	viper.Viper()
	initmysql.PingMySql(viper.ServerConfig, viper.DBConfig)
	//initRedis.ConnectRedis(viper.ServerConfig)
	initmysql.AutoMigrateTable()
	engine := gin.New()
	initrouter.LoadRouter(engine)
	startServer(engine)
}

func startServer(engine *gin.Engine) {
	// 设置日志写入器为文件
	file, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = file
	engine.Run(":9999")
}
