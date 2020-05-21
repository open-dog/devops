package manage

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReleaseInfo(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "info.html", nil)
}