package release

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"go-admin/app/helper"
	"go-admin/app/model/devops/releasestruct"
	"go-admin/app/service/release"
)

// 添加工单
func AddRelease(ctx *gin.Context) {
	//解析参数
	byte_param := helper.GetContent(ctx)

	add_param := new(releasestruct.AddReleaseParam)
	if err:= json.Unmarshal(byte_param, add_param); err != nil {
		helper.ErrorResponse(ctx, nil, err.Error())
		return
	}

	err := release.NewReleaseService().Add(add_param)

	if err != nil {
		helper.ErrorResponse(ctx, nil, err.Error())
		return
	}
	helper.SuccessResponse(ctx, g.List{})


}
