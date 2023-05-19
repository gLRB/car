package main

import (
	"backend/app/model"
	. "backend/app/services"
	"github.com/gin-gonic/gin"
)

func main() {

	model.MysqlInit()

	router := gin.Default()
	router.GET("/test", Test)
	router.POST("/new_car", New)
	router.GET("/show", Show)
	router.POST("/edit", Edit)
	_ = router.Run()

}
