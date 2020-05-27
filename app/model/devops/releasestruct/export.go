package releasestruct

type OnlineInfo struct {
	Id      int              `json:"id"`
	Content string           `json:"content"`
	Author  string           `json:"author"`
	Sql     []*OnlineSql     `json:"sql"`
	Package []*OnlinePackage `json:"package"`
	Service []*OnlineService `json:"service"`
}

func NewOnlineInfo() *OnlineInfo {
	return &OnlineInfo{}
}

type OnlineSql struct {
	Db     string `json:"db"`
	Crud   string `json:"crud"`
	Remark string `json:"remark"`
}

type OnlinePackage struct {
	PackageName string `json:"package_name"`
	BranchName  string `json:"branch_name"`
}

type OnlineService struct {
	ServiceName string    `json:"service_name"`
	Version     string    `json:"version"`
	BranchName  string    `json:"branch_name"`
	Env         []*Env    `json:"env"`
	Script      []*Script `json:"script"`
}
