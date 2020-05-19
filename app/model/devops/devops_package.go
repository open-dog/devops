package devops

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetDevopsPackageTable(ctx *context.Context) table.Table {

	devopsPackageTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := devopsPackageTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Release_id", "release_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Branch", "branch", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("devops_package").SetTitle("Devops_package").SetDescription("Devops_package")

	formList := devopsPackageTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Release_id", "release_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Branch", "branch", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("devops_package").SetTitle("Devops_package").SetDescription("Devops_package")

	return devopsPackageTable
}
