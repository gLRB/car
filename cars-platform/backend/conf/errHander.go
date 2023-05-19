package conf

import "github.com/gin-gonic/gin"

func ErrRet(msg string) gin.H {
	return gin.H{
		"code":    10000,
		"message": msg,
		"data":    map[string]string{},
	}
}
