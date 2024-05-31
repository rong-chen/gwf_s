package user

import (
	"gwf/global"
)

type Info struct {
	Name     string `gorm:"column:name" json:"name" comment:"用户姓名" validate:"required"`
	Sex      int    `gorm:"column:sex" json:"sex" comment:"用户性别" validate:"required" `
	Phone    string `gorm:"column:phone;unique" json:"phone" comment:"手机号码" validate:"required"`
	Email    string `gorm:"column:email;unique" json:"email" comment:"邮箱" validate:"required"`
	UserName string `gorm:"column:username;" json:"username" comment:"用户名" validate:"required"`
	Password string `gorm:"column:password" json:"-" comment:"密码(密文)" validate:"required"`
	Role     string `gorm:"column:role" json:"role" comment:"用户角色(admin:最高权限,editUser:编辑员,user:普通成员)"`
	global.Model
}
type Params struct {
	Info
	Code string `json:"code" form:"code"`
}
type Login struct {
	Email    string `json:"email" gorm:"email"  validate:"required"`
	Password string `json:"password" gorm:"password"  validate:"required"`
}

func (Info) TableName() string {
	return "user"
}
