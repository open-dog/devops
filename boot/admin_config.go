package boot

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var (
	cfg config.Config
)

//初始化对应的变量
func init()  {
	// GoAdmin全局配置，也可以写成一个json，通过json引入
	cfg = config.Config{
		// 数据库配置，为一个map，key为连接名，value为对应连接信息
		Databases: GetDatabaseList(),
		UrlPrefix: "admin", // 访问网站的前缀
		// Store 必须设置且保证有写权限，否则增加不了新的管理员用户
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		// 网站语言
		Language: language.CN,
	}
}

func GetDatabaseList() config.DatabaseList  {
	db := g.Cfg().Get("database.default.0")
	return config.DatabaseList{
		// 默认数据库连接，名字必须为default
		"default": {
			Host:       gconv.String(db.(map[string]interface{})["host"]),
			Port:       gconv.String(db.(map[string]interface{})["port"]),
			User:       gconv.String(db.(map[string]interface{})["user"]),
			Pwd:        gconv.String(db.(map[string]interface{})["pass"]),
			Name:       gconv.String(db.(map[string]interface{})["name"]),
			MaxIdleCon: 50,
			MaxOpenCon: 150,
			Driver:     config.DriverMysql,
		},
	}
}

// 返回对应的配置文件
func Cfg() config.Config  {
	return cfg
}
