package release

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"go-admin/app/helper"
	"go-admin/app/service/release"
)

//详情
func InfoRelease(ctx *gin.Context) {

	release_id := gconv.Int(ctx.Query("release_id"))
	res, err := release.NewReleaseService().Info(gconv.Int(release_id))
	if err != nil {
		helper.ErrorResponse(ctx, nil, err.Error())
	} else {
		helper.SuccessResponse(ctx, res)
	}
}
