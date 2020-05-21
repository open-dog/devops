package release

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"go-admin/app/helper"
	"go-admin/app/service/haijigit"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// 分支相关操作
func Branch(ctx *gin.Context)  {
	name_param := ctx.Query("name")

	//解析是否多个分支一起请求
	name_slice := strings.Split(name_param, ",")

	//接收对应的分支数据
	branch_list_ch := make(chan g.Map, len(name_slice))

	//并发请求
	wg.Add(len(name_slice))
	for _, name := range name_slice{
		go func(name string, ch *chan g.Map) {
			defer wg.Done()
			branch_list, err := haijigit.GetBranchNameByName(name)
			if err == nil {
				res := g.Map{
					name : branch_list,
				}
				branch_list_ch <- res
			}
		}(name, &branch_list_ch)
	}
	wg.Wait()
	close(branch_list_ch)
	res := g.Map{}
	for item := range branch_list_ch {
		for k, v := range item{
			res[k] = v
		}
	}
	helper.SuccessResponse(ctx, res)

}
