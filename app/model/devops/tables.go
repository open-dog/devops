package devops

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "devops_package" => http://localhost:9033/admin/info/devops_package
// "devops_release" => http://localhost:9033/admin/info/devops_release
// "devops_service" => http://localhost:9033/admin/info/devops_service
// "devops_sql" => http://localhost:9033/admin/info/devops_sql
//
// example end
//
var Generators = map[string]table.Generator{
	"devops_package": GetDevopsPackageTable,
	"devops_release": GetDevopsReleaseTable,
	"devops_service": GetDevopsServiceTable,
	"devops_sql":     GetDevopsSqlTable,

	// generators end
}
