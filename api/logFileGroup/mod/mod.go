package mod

import (
	user "gwf/api/user/mod"
	"gwf/global"
)

type LogFileGroup struct {
	global.Model
	Title        string    `gorm:"column:title" json:"title" comment:"标题" validate:"required"`
	Content      string    `gorm:"column:content" json:"content" comment:"备注或者文件内容" validate:"required"`
	Type         string    `gorm:"column:type" json:"type" comment:"类型：文件file or 文件夹dir" validate:"required"`
	CreateUserID string    `gorm:"column:create_user_id;type:char(36);index;comment:创建人ID" json:"createUserId"`
	ParseDirID   string    `gorm:"column:parse_dir_id" json:"parseDirID" comment:"父文件夹ID"`
	CreateUser   user.Info `gorm:"foreignKey:CreateUserID;references:ID" json:"create_user" comment:"创建人"`
}
