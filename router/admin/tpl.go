package admin

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"go-admin/app/model/pages"
)

func InitRouter(eng *engine.Engine) {

	eng.HTML("GET", "/admin/form", pages.GetFormContent)
	eng.HTML("GET", "/admin/table", pages.GetTableContent)


	//添加工单
	eng.HTMLFile("GET", "/admin/devops/release/add", "./public/html/release/add.tpl", map[string]interface{}{
		"msg": "Hello world",
	})

	//编辑工单
	eng.HTMLFile("GET", "/admin/devops/release/edit", "./public/html/release/edit.tpl", map[string]interface{}{
		"msg": "Hello world",
	})

	//工单列表
	eng.HTMLFile("GET", "/admin/devops/release/list", "./public/html/release/list.tpl", map[string]interface{}{
		"msg": "Hello world",
	})



}