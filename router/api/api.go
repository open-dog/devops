package api

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/api/v1/devops/release"
)

func InitRouter(r *gin.Engine)  {
	apiv1 := r.Group("/v1")
	{
		devops := apiv1.Group("/devops")
		{
			devops.POST("/release/add", release.AddRelease)
			devops.GET("/release/info", release.InfoRelease)
			devops.GET("/release/list", release.ListRelease)
			devops.POST("/release/edit", release.EditRelease)
			devops.GET("/git/branch", release.Branch)
		}
	}
}