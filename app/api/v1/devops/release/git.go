package release

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/helper"
	"go-admin/app/service/haijigit"
)

// 分支相关操作
func Branch(ctx *gin.Context)  {
	name := ctx.Query("name")

	branch_list, err := haijigit.GetBranchNameByName(name)
	if err != nil {
		helper.ErrorResponse(ctx, nil, err.Error())
	} else {
		helper.SuccessResponse(ctx, branch_list)
	}
}
