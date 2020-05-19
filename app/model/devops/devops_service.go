package devops

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetDevopsServiceTable(ctx *context.Context) table.Table {

	devopsServiceTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := devopsServiceTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Release_id", "release_id", db.Int)
	info.AddField("Service_name", "service_name", db.Varchar)
	info.AddField("Branch", "branch", db.Varchar)
	info.AddField("Env", "env", db.Text)
	info.AddField("Script", "script", db.Text)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("devops_service").SetTitle("Devops_service").SetDescription("Devops_service")

	formList := devopsServiceTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Release_id", "release_id", db.Int, form.Number)
	formList.AddField("Service_name", "service_name", db.Varchar, form.Text)
	formList.AddField("Branch", "branch", db.Varchar, form.Text)
	formList.AddField("Env", "env", db.Text, form.Text)
	formList.AddField("Script", "script", db.Text, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("devops_service").SetTitle("Devops_service").SetDescription("Devops_service")

	return devopsServiceTable
}
