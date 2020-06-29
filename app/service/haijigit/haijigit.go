package haijigit

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"go-admin/app/helper"
	"strings"
)

//海际git相关的操作

//配置相关的项目

// 返回服务相关的信息
func ServerNameList() g.List {
	return g.List{
		{"key": "goods-server", "val": "git@code.aliyun.com:haiji-server/goods-server.git"},
		{"key": "order-server", "val": "git@code.aliyun.com:haiji-server/order-server.git"},
		{"key": "pay-server", "val": "git@code.aliyun.com:haiji-server/pay-server.git"},
		{"key": "user-server", "val": "git@code.aliyun.com:haiji-server/user-server.git"},
		{"key": "xianglong-server", "val": "git@code.aliyun.com:haiji-server/xianglong-server.git"},
		{"key": "idc-server", "val": "git@code.aliyun.com:haiji-server/idc-server.git"},
		{"key": "stores-server", "val": "git@code.aliyun.com:haiji-server/stores-server.git"},
		{"key": "api-server", "val": "git@code.aliyun.com:haiji-server/api-server.git"},
		{"key": "tapin-server", "val": "git@code.aliyun.com:haiji-server/tapin-server.git"},
		{"key": "manage-server", "val": "git@code.aliyun.com:haiji-server/manage-server.git"},
		{"key": "salesman-server", "val": "git@code.aliyun.com:haiji-server/salesman-server.git"},
		{"key": "upload-server", "val": "git@code.aliyun.com:haiji-server/upload-server.git"},
		{"key": "sms-server", "val": "git@code.aliyun.com:haiji-server/sms-server.git"},
	}
}

//返回package相关的项目
func PackageNameList() g.List {
	return g.List{
		{"key": "goods-package", "val": "git@code.aliyun.com:haiji-package/goods-package.git"},
		{"key": "order-package", "val": "git@code.aliyun.com:haiji-package/order-package.git"},
		{"key": "user-package", "val": "git@code.aliyun.com:haiji-package/user-package.git"},
		{"key": "pay-package", "val": "git@code.aliyun.com:haiji-package/pay-package.git"},
		{"key": "xianglong-package", "val": "git@code.aliyun.com:haiji-package/xianglong-package.git"},
		{"key": "idc-package", "val": "git@code.aliyun.com:haiji-package/idc-package.git"},
		{"key": "stores-package", "val": "git@code.aliyun.com:haiji-package/stores-package.git"},
		{"key": "sms-package", "val": "git@code.aliyun.com:haiji-package/sms-package.git"},
		{"key": "dingtalk-center", "val": "git@code.aliyun.com:haiji-package/dingtalk-center.git"},
		{"key": "common-config", "val": "git@code.aliyun.com:haiji-test/common-config.git"},
		{"key": "haiji-apicenter", "val": "git@code.aliyun.com:haiji-package/haiji-apicenter.git"},
		{"key": "epet-center", "val": "git@code.aliyun.com:haiji-package/epet-center.git"},
	}
}

func GetServerAddress(server_name string) string {
	for _, service := range ServerNameList() {
		if service["key"] == server_name {
			if val, ok := service["val"]; ok {
				return gconv.String(val)
			}
		}
	}
	return ""
}

func GetPackageAddress(package_name string) string {
	for _, p := range PackageNameList() {
		if p["key"] == package_name {
			if val, ok := p["val"]; ok {
				return gconv.String(val)
			}
		}
	}
	return ""
}

func GetBranchNameByName(name string) (g.List, error) {
	branch_list := make(g.List, 0)
	code_address := GetServerAddress(name)
	if code_address == "" {
		code_address = GetPackageAddress(name)
	}
	if code_address == "" {
		return branch_list, errors.New("对应的代码仓库不存在")
	}
	fmt.Println(code_address)

	cmd := "cd ./code/" + name + " && git remote update -p && git branch -r"
	out, err := helper.Cmd(cmd, true)
	if err != nil {
		return branch_list, err
	}
	list_arr := strings.Split(string(out), "\n")

	for _, item := range list_arr {
		if item != "" && !strings.Contains(item, "HEAD") && !strings.Contains(item, "正在") {
			branch_list = append(branch_list, g.Map{"label": strings.TrimSpace(item), "value": strings.TrimSpace(item)})
		}
	}

	return branch_list, nil
}
