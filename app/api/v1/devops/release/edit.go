package release

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"go-admin/app/helper"
	"go-admin/app/model/devops/releasestruct"
	"go-admin/app/service/release"
	"log"
)

//编辑
func EditRelease(ctx *gin.Context)  {
	byte_param := helper.GetContent(ctx)

	edit_param := new(releasestruct.EditReleaseParam)
	if err:= json.Unmarshal(byte_param, edit_param); err != nil {
		log.Panic(err)
	}
	err := release.NewReleaseService().Edit(edit_param)

	if err != nil {
		log.Fatal(err)
	}
	helper.SuccessResponse(ctx, g.List{})

}
