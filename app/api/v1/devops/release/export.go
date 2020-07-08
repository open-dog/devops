package release

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-admin/app/helper"
	"go-admin/app/service/release"
	"io"
)

func ExportRelease(ctx *gin.Context) {

	release_id := ctx.Query("release_id")

	resp, err := release.NewReleaseService().Export(release_id)
	if err != nil {
		helper.ErrorResponse(ctx, nil, err.Error())
		return
	}
	b_str, err := json.Marshal(resp)
	if err != nil {
		helper.ErrorResponse(ctx, nil, err.Error())
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename=online"+release_id+".json")
	ctx.Header("Content-Type", ctx.GetHeader("Content-Type"))
	io.Copy(ctx.Writer, bytes.NewReader(b_str))
}
