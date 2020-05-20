package release

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/helper"
	"go-admin/app/service/common"
)

//返回一些通用信息
//数据库名， 包名， 服务名
func GetCommon(ctx *gin.Context){
	cfg := common.GetCommonCfg()

	helper.SuccessResponse(ctx, cfg)
}
