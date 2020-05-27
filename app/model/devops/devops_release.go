package devops

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

const (
	STATUS_INI    = 1 //提交中
	STATUS_PRE    = 2 //预发布(不能修改)
	STATUS_ONLINE = 3 //已发布
	STATUS_DEL    = 8 //删除
)

func StatusList() g.MapIntStr {
	return g.MapIntStr{
		STATUS_INI:    "待发布",
		STATUS_PRE:    "预发布",
		STATUS_ONLINE: "已发布",
		STATUS_DEL:    "已删除",
	}
}

func GetDevopsReleaseTable(ctx *context.Context) table.Table {

	//初始化数据模型
	devopsReleaseTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := devopsReleaseTable.GetInfo()

	info.AddField("工单Id", "id", db.Int).FieldFilterable()
	info.AddField("工单标题", "title", db.Varchar).FieldFilterable()
	info.AddField("提交人", "author", db.Varchar)
	info.AddField("工单内容", "content", db.Varchar)
	// FieldDisplay 根据字段输出对应的内容
	info.AddField("状态", "status", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		return StatusList()[gconv.Int(model.Value)]
	}).FieldFilterable(types.FilterType{FormType: form.Select}).
		FieldFilterOptions(types.FieldOptions{
			//todo 需要整理
			{Value: "1", Text: "待发布"},
			{Value: "2", Text: "预发布"},
			{Value: "3", Text: "已发布"},
			{Value: "8", Text: "已删除"},
		}).FieldFilterOptionExt(map[string]interface{}{"allowClear": true})

	info.AddField("添加时间", "created_at", db.Timestamp)
	info.AddField("修改时间", "upated_at", db.Timestamp)

	// 设置页面标题和描述
	info.SetTable("devops_release").SetTitle("工单列表").SetDescription("工单列表")

	//禁用创建按钮
	info.HideNewButton()
	//隐藏编辑按钮
	//info.HideEditButton()
	//info.HideDeleteButton()

	// 设置隐藏搜索框
	info.HideFilterArea()

	//设置筛选表单的宽度
	info.SetFilterFormHeadWidth(1)
	info.SetFilterFormInputWidth(2)

	//新增按钮
	info.AddButton("添加工单", icon.Adn, action.Jump("devops_release_add"))

	//增加列按钮操作
	info.AddActionButton(`<i class="fa fa-edit"></i>`, action.Jump("/admin/info/permission/edit"))

	formList := devopsReleaseTable.GetForm()

	formList.AddField("工单Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("工单标题", "title", db.Varchar, form.Text)
	formList.AddField("工单内容", "content", db.Varchar, form.Text)

	// FieldHide 隐藏表单相关的值
	formList.AddField("Author", "author", db.Varchar, form.Text).FieldHide()
	formList.AddField("Status", "status", db.Tinyint, form.Number).FieldHide()
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime).FieldHide()
	formList.AddField("Upated_at", "upated_at", db.Timestamp, form.Datetime).FieldHide()

	formList.SetTable("devops_release").SetTitle("添加工单").SetDescription("添加工单")

	return devopsReleaseTable
}
