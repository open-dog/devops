package main

import (
	"context"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // 引入适配器，必须引入，如若不引入，则需要自己定义
	"github.com/GoAdminGroup/go-admin/engine"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // 引入对应数据库引擎
	_ "github.com/GoAdminGroup/go-admin/plugins/admin"
	_ "github.com/GoAdminGroup/themes/adminlte" // 引入主题，必须引入，不然报错
	"go-admin/app/model/devops"
	"go-admin/boot"
	"go-admin/router"
	"go-admin/router/admin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	r := router.Router()

	// 实例化一个GoAdmin引擎对象
	eng := engine.Default()

	// 增加配置与插件，使用Use方法挂载到Web框架中
	_ = eng.AddConfig(boot.Cfg()).
		// 这里引入你需要管理的业务表配置
		// 后面会介绍如何使用命令行根据你自己的业务表生成Generators
		AddGenerators(devops.Generators).
		Use(r)

	//注册一些web页面到路由中
	admin.InitRouter(eng)

	//实现服务的优雅关闭与重启
	srv := &http.Server{
		Addr:    ":8089",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit //阻塞，直到接收到信号

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown: ", err)
	}
}
