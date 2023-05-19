package model

import (
	. "backend/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
	使用docker部署MySQL，需要注意设置目录映射关系 （宿主机目录:docker container目录）
	docker run -p 3306:3306
	--privileged=true
	-v /Users/alan/Desktop/computer/docker/mysql/log:/var/log/mysql
	-v /Users/alan/Desktop/computer/docker/mysql/data:/var/lib/mysql
	-v /Users/alan/Desktop/computer/docker/mysql/conf/:/etc/mysql/conf.d
	-e MYSQL_ROOT_PASSWORD=123456
	--name root -d
	mysql

docker run -p 3306:3306 --privileged=true -v /Users/alan/Desktop/computer/docker/mysql/log:/var/log/mysql -v /Users/alan/Desktop/computer/docker/mysql/data:/var/lib/mysql -v /Users/alan/Desktop/computer/docker/mysql/conf/:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=123456 --name root -d mysql

*/

func NewDBObject() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(GetDBConf()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
