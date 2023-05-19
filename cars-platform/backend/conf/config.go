package conf

import "fmt"

func GetCurEnv() string {
	return "prod"
}

func GetDBConf() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", LocalUserName, LocalPassWord, LocalHost, LocalPort, DBName, Timeout)
	//fmt.Print("dsn", dsn)
	return dsn
}
