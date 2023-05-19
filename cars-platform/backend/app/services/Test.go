package services

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Test(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg":  "insert data fail",
		"data": "",
	})
	log.Print("hello")
}
