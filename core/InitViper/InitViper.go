package viper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	ServerConfig Server
	DBConfig     DB
)

type Server struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type DB struct {
	Database string `json:"database"`
}

func Viper() {
	environment := os.Getenv("ENV")
	fmt.Println(environment)
	if environment == "" {
		environment = "development"
	}
	if environment == "development" {
		viper.SetConfigFile("config.dev.yaml")
		viper.AddConfigPath("../../")
	} else if environment == "production" {
		viper.SetConfigFile("config.yaml")
		viper.AddConfigPath("./")
	} else {
		fmt.Println("未知环境")
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Error reading config file: %s\n", err))
	}
	ServerConfig.Host = viper.GetString("server.host")
	ServerConfig.Port = viper.GetInt("server.port")
	ServerConfig.Username = viper.GetString("server.username")
	ServerConfig.Password = viper.GetString("server.password")
	DBConfig.Database = viper.GetString("mysql.database")
}
