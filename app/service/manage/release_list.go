package manage

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReleaseList(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
