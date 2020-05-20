package common

import (
	"github.com/gogf/gf/frame/g"
	"go-admin/app/service/haijigit"
)

var db_slice = []string{
	"haiji_activity",
	"haiji_admin",
	"haiji_etl",
	"haiji_goods",
	"haiji_idc",
	"haiji_message",
	"haiji_order",
	"haiji_salesman",
	"haiji_tapin",
	"haiji_user",
	"new_haiji_admin",
}

// 返回数据的列表显示
func GetDbList() g.List {
	db_list := make(g.List, 0)
	for _,db_name := range db_slice  {
		db_list = append(db_list, g.Map{"label":db_name, "value":db_name})
	}
	return db_list
}

// 返回package数据
func GetPackageList() g.List {
	p_list := make(g.List, 0)
	p_name_list :=  haijigit.PackageNameList()
	for _, p := range p_name_list {
		if p_name, ok := p["key"]; ok {
			p_list = append(p_list, g.Map{"label":p_name, "value":p_name})
		}
	}
	return p_list
}

// 返回service数据
func GetServiceList() g.List {
	s_list := make(g.List, 0)
	s_name_list :=  haijigit.ServerNameList()
	for _, p := range s_name_list {
		if p_name, ok := p["key"]; ok {
			s_list = append(s_list, g.Map{"label":p_name, "value":p_name})
		}
	}
	return s_list
}

//返回公共配置
func GetCommonCfg() g.Map {
	return g.Map{
		"databases" : GetDbList(),
		"packages" : GetPackageList(),
		"services" : GetServiceList(),
	}
}