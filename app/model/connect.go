package model

import (
	"database/sql"
	"log"
)

// 暂时不用，配置解析太麻烦了
func Db() *sql.DB {
	driverName := "mysql"
	dataSourceName := "root:root@tcp(127.0.0.1:3306)/go_admin"

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
