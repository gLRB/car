//package main
//
//import (
//	"backend/app/model"
//	. "backend/app/services"
//	"crypto/tls"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//)
//
//func main() {
//
//	model.MysqlInit()
//
//	router := gin.Default()
//	router.GET("/test", Test)
//	router.POST("/new_car", New)
//	router.GET("/show", Show)
//	router.POST("/edit", Edit)
//	_ = router.Run("0.0.0.0:8003")
//
//}

package main

import (
	. "backend/app/services"
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 加载 SSL 证书和私钥文件
	cert, err := tls.LoadX509KeyPair("./conf/cert.pem", "./conf/key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// 创建 TLS 配置
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	//model.MysqlInit()

	router := gin.Default()
	router.GET("/test", Test)
	router.POST("/new_car", New)
	router.GET("/show", Show)
	router.POST("/edit", Edit)

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:      ":8003", // 监听 8443 端口
		Handler:   router,  // Gin 路由器
		TLSConfig: config,  // TLS 配置
	}

	// 启动 HTTPS 服务器
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}
}
