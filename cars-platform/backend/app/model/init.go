package model

import (
	"backend/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlInit() {
	dsn := conf.GetDBConf()
	dbd, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	err = dbd.AutoMigrate(
		&Cars{},
	)
	if err != nil {
		panic("迁移数据库失败, error=" + err.Error())
	}
}
