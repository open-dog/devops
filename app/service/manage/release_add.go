package manage

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReleaseAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add.html", nil)
}
