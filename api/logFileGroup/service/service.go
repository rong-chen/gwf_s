package service

import (
	"fmt"
	logFileGroup "gwf/api/logFileGroup/mod"
	"gwf/global"
)

func CreateRow(data *logFileGroup.LogFileGroup) error {
	return global.DB.Create(&data).Error
}

func SearchDir(id string, cid string) (list []logFileGroup.LogFileGroup) {
	fmt.Println(id)
	global.DB.Where("parse_dir_id = ? and create_user_id = ?", id, cid).Find(&list)
	return
}
