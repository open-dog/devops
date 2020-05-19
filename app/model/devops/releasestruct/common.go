package releasestruct


//添加工单的请求参数
type AddReleaseParam struct {
	Title   string     `json:"title" form:"title"`  //解析对应的工单名
	Author  string     `json:"author" form:"title"` //解析对应的提交者名
	Sql     *[]Sql     `json:"sql"`                 //解析对应需要执行的sql语句
	Package *[]Package `json:"package"`             //解析对应的包名
	Service *[]Service  `json:"service"`             //解析对应脚本名

}

type Service struct {
	Name   string `json:"name"`
	Branch string `json:"branch"`
	Env []*Env `json:"env"`
	Script []*Script `json:"script"`
}

type Env struct {
	Handle string `json:"handle"`
	Key    string `json:"key"`
	// 除了是字符串类型还可能是其他类型
	Value interface{} `json:"value"`
}

type Script struct {
	Cmd    string `json:"cmd"`
	Remark string `json:"remark"`
}

type Package struct {
	Name   string `json:"name"`
	Branch string `json:"branch"`
}

type Sql struct {
	Db     string `json:"db"`
	Crud   string `json:"crud"`
	Remark string `json:"remark"`
}