package release

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"go-admin/app/helper"
	"go-admin/app/model/devops/releasestruct"
	"go-admin/app/service/release"
)

//编辑
func EditRelease(ctx *gin.Context) {
	byte_param := helper.GetContent(ctx)

	edit_param := new(releasestruct.EditReleaseParam)
	if err := json.Unmarshal(byte_param, edit_param); err != nil {
		helper.ErrorResponse(ctx, nil, err.Error())
		return
	}
	err := release.NewReleaseService().Edit(edit_param)

	if err != nil {
		helper.ErrorResponse(ctx, nil, err.Error())
		return
	}
	helper.SuccessResponse(ctx, g.List{})

}
