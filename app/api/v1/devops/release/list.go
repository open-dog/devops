package release

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"go-admin/app/helper"
	"go-admin/app/service/release"
)

//列表
func ListRelease(ctx *gin.Context)  {
	release_id := gconv.Int(ctx.Query("id"))
	author := ctx.Query("author")
	page := gconv.Int(ctx.Query("page"))

	list, err := release.NewReleaseService().List(release_id, author, page)
	if err != nil {
		helper.ErrorResponse(ctx, nil, err.Error())
	} else{
		helper.SuccessResponse(ctx, list)
	}
}
