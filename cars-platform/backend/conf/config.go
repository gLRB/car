package conf

import "fmt"

func GetCurEnv() string {
	return "local"
}

func GetDBConf() string {
	if GetCurEnv() == "local" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", LocalUserName, LocalPassWord, LocalHost, LocalPort, DBName, Timeout)
		//fmt.Print("dsn", dsn)
		return dsn
	} else {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", ProdUserName, ProdPassWord, ProdHost, ProdPort, DBName, Timeout)
		//fmt.Print("dsn", dsn)
		return dsn
	}
}
