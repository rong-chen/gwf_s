package global

import (
	"github.com/go-redis/redis"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

type Model struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func GinBackData(code int, msg string, data interface{}) *gin.H {
	return &gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
