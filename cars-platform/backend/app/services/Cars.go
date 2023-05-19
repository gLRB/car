package services

import (
	"backend/app/model"
	"backend/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New(ctx *gin.Context) {
	// new one car message
	var req model.NewReq
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Println("err -> ", err)
		ctx.JSON(http.StatusOK, conf.ErrRet("解析入参失败"))
		return
	}

	if !CheckToken(req.Token) {
		ctx.JSON(http.StatusOK, conf.ErrRet("凭证校验失败"))
		return
	}

	fmt.Printf("入参 -> %v\n", req)
	db, err := model.NewDBObject()
	if err != nil {
		ctx.JSON(http.StatusOK, conf.ErrRet("数据库连接失败"))
		return
	}
	if err := model.InsertCarMsg(db, req); err != nil {
		fmt.Printf("InsertCarMsg err -> %v", err)
		ctx.JSON(http.StatusOK, conf.ErrRet("插入车辆信息到数据库失败"))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    map[string]string{},
	})

}

func Show(ctx *gin.Context) {
	// return cars msg
	var req model.ShowReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, conf.ErrRet("Show解析入参失败"))
		return
	}
	db, err := model.NewDBObject()
	if err != nil {
		ctx.JSON(http.StatusOK, conf.ErrRet("数据库连接失败"))
		return
	}
	cars, err := model.SearchCars(db, req)
	if err != nil {
		ctx.JSON(http.StatusOK, conf.ErrRet("查找数据失败"))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    cars,
	})

}

func Edit(ctx *gin.Context) {

	var req model.EditReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, conf.ErrRet("Edit解析入参失败"))
		return
	}

	if !CheckToken(req.Token) {
		ctx.JSON(http.StatusOK, conf.ErrRet("凭证校验失败"))
		return
	}

	db, err := model.NewDBObject()
	if err != nil {
		ctx.JSON(http.StatusOK, conf.ErrRet("数据库连接失败"))
		return
	}
	fmt.Println("EditReq", req)
	if err := model.EditCar(db, req); err != nil {
		ctx.JSON(http.StatusOK, conf.ErrRet("编辑车辆信息失败"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    req,
	})

}
