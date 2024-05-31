package initmysql

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm/logger"
	logFileGroup "gwf/api/logFileGroup/mod"
	user "gwf/api/user/mod"
	viper "gwf/core/InitViper"
	"gwf/global"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 存储已有的结构体数据库表
var tablesArray = []interface{}{
	&user.Info{},
	&logFileGroup.LogFileGroup{},
}

func PingMySql(serverConfig viper.Server, dbConfig viper.DB) {
	// 连接到MySQL数据库
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", serverConfig.Username, serverConfig.Password, serverConfig.Host, serverConfig.Port))
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()
	err = db.QueryRow("SELECT SCHEMA_NAME FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = " + fmt.Sprintf("'%s'", dbConfig.Database)).Scan(&dbConfig.Database)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			_, err2 := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbConfig.Database)
			if err2 != nil {
				fmt.Println("创建数据库失败:", err)
				return
			}
		} else {
			fmt.Println("查询失败:", err)
			return
		}
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", serverConfig.Username, serverConfig.Password, serverConfig.Host, serverConfig.Port, dbConfig.Database)

	// 配置 Logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略 ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 是否启用彩色打印
		},
	)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		fmt.Println("初始化数据库失败")
		return
	}
}

func AutoMigrateTable() {
	for i := 0; i < len(tablesArray); i++ {
		global.DB.AutoMigrate(tablesArray[i])
	}
}
