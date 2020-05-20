package db

import "github.com/gogf/gf/frame/g"

func GetDbList() g.List {
	db_list := g.List{
		{"label":"haiji_activity", "value":"haiji_activity"},
		{"label":"haiji_admin", "value":"haiji_admin"},
		{"label":"haiji_goods", "value":"haiji_activity"},
		{"label":"haiji_activity", "value":"haiji_activity"},
		{"label":"haiji_activity", "value":"haiji_activity"},
		{"label":"haiji_activity", "value":"haiji_activity"},
	}

	return db_list
}