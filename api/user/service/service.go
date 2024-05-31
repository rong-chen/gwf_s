package service

import (
	"fmt"
	user "gwf/api/user/mod"
	"gwf/global"
	"gwf/utils"
)

func Create(user *user.Info) error {
	err := global.DB.Create(&user).Error
	return err
}

func Find(key, phone string) (*user.Info, error) {
	var user user.Info
	err := global.DB.Where(fmt.Sprintf("%s = ?", key), phone).First(&user).Error
	return &user, err
}

func Login(email string, password string) (Info *user.Info, msg string) {
	err := global.DB.Where("email = ?", email).First(&Info).Error
	if err != nil {
		return nil, "暂无该用户"
	}
	ok := utils.ValidPassword(password, Info.Password)
	if ok {
		return Info, "登录成功"
	} else {
		return nil, "密码错误"
	}
}
