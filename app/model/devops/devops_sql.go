package devops

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetDevopsSqlTable(ctx *context.Context) table.Table {

	devopsSqlTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := devopsSqlTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Release_id", "release_id", db.Int)
	info.AddField("Db", "db", db.Varchar)
	info.AddField("Crud", "crud", db.Varchar)
	info.AddField("Status", "status", db.Tinyint)
	info.AddField("Remark", "remark", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("devops_sql").SetTitle("Devops_sql").SetDescription("Devops_sql")

	formList := devopsSqlTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Release_id", "release_id", db.Int, form.Number)
	formList.AddField("Db", "db", db.Varchar, form.Text)
	formList.AddField("Crud", "crud", db.Varchar, form.Text)
	formList.AddField("Status", "status", db.Tinyint, form.Number)
	formList.AddField("Remark", "remark", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("devops_sql").SetTitle("Devops_sql").SetDescription("Devops_sql")

	return devopsSqlTable
}
