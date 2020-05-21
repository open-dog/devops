package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/helper"
)

//错误处理中间件
func Err() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 处理请求
		c.Next()

		defer func() {
			if err := recover(); err != nil {
				helper.ErrorResponse(c,nil, "请求失败")
			}
		}()

		
	}
}
