package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/router/api"
)

var (
	router *gin.Engine
)

func init()  {
	router = gin.Default()
	//注册api相关路由
	api.InitRouter(router)
}

func Router() *gin.Engine {
	return router
}

