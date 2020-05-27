package release

import (
	"encoding/json"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"go-admin/app/model/devops"
	"go-admin/app/model/devops/releasestruct"
)

//添加工单操作service

type ReleaseService struct {
	Db gdb.DB
}

const (
	PERPAGE = 20
)

func NewReleaseService() *ReleaseService {
	return &ReleaseService{Db: g.DB()}
}

func (r *ReleaseService) Add(data *releasestruct.AddReleaseParam) error {

	//开启事务
	if tx, err := r.Db.Begin(); err == nil {
		//写入主干数据
		release_res, err := tx.Table("devops_release").Data(g.Map{
			"title":      data.Title,
			"author":     data.Author,
			"content":    data.Content,
			"status":     1,
			"created_at": gtime.Datetime(),
			"updated_at": gtime.Datetime(),
		}).Save()
		if err != nil {
			tx.Rollback()
			return err
		}
		release_id, err := release_res.LastInsertId()
		if err != nil {
			return err
		}

		//写入sql表
		sql_list := make([]g.Map, 0)
		for _, s := range *data.Sql {
			sql_list = append(sql_list, g.Map{
				"release_id": release_id,
				"db":         s.Db,
				"crud":       s.Crud,
				"remark":     s.Remark,
				"status":     1,
				"created_at": gtime.Datetime(),
				"updated_at": gtime.Datetime(),
			})
		}
		_, err = tx.Table("devops_sql").Data(sql_list).Save()
		if err != nil {
			tx.Rollback()
			return err
		}
		//写入package
		package_list := make([]g.Map, 0)
		for _, p := range *data.Package {
			package_list = append(package_list, g.Map{
				"release_id": release_id,
				"name":       p.Name,
				"branch":     p.Branch,
				"created_at": gtime.Datetime(),
				"updated_at": gtime.Datetime(),
			})
		}
		_, err = tx.Table("devops_package").Data(package_list).Save()
		if err != nil {
			tx.Rollback()
			return err
		}

		//写入service
		service_list := make([]g.Map, 0)
		for _, s := range *data.Service {
			service_list = append(service_list, g.Map{
				"release_id": release_id,
				"name":       s.Name,
				"branch":     s.Branch,
				"version":    s.Version,
				"env":        gconv.String(s.Env),
				"script":     gconv.String(s.Script),
				"created_at": gtime.Datetime(),
				"updated_at": gtime.Datetime(),
			})
		}
		_, err = tx.Table("devops_service").Data(service_list).Save()
		if err != nil {
			tx.Rollback()
			return err
		}

		tx.Commit()
	}
	return nil
}

func (r *ReleaseService) Info(release_id int) (g.Map, error) {

	res := g.Map{}
	//获取主表数据
	info, err := r.Db.Table("devops_release").Where("id=?", release_id).One()
	if err != nil {
		return res, err
	}
	res = info.Map()

	//获取sql表数据
	sql, err := r.Db.Table("devops_sql").Where("release_id=?", release_id).All()
	if err != nil {
		return res, err
	}
	res["sql"] = sql.List()

	//获取package数据
	pg, err := r.Db.Table("devops_package").Where("release_id=?", release_id).All()
	if err != nil {
		return res, err
	}
	res["package"] = pg.List()

	//获取service的数据
	service, err := r.Db.Table("devops_service").Where("release_id=?", release_id).All()
	if err != nil {
		return res, err
	}

	service_list := service.List()
	for _, s := range service_list {
		env := g.List{}
		script := g.List{}
		json.Unmarshal([]byte(gconv.String(s["env"])), &env)
		json.Unmarshal([]byte(gconv.String(s["script"])), &script)
		s["env"] = env
		s["script"] = script
	}
	res["service"] = service_list

	return res, nil
}

//列表
func (r *ReleaseService) List(release_id int, author string, page int) (g.Map, error) {

	res := make(g.Map)
	//对page页面做处理
	if g.IsEmpty(page) {
		page = 1
	}

	//解析参数
	mod := r.Db.Table("devops_release")
	if !g.IsEmpty(release_id) {
		mod = mod.Where("id = ?", release_id)
	}
	if !g.IsEmpty(author) {
		mod = mod.Where("author like ?", "%"+author+"%")
	}

	//查询总页数
	total_num, err := mod.Count()
	if err != nil {
		return res, err
	}

	//处理分页
	if g.IsEmpty(page) {
		min_limit := (page - 1) * PERPAGE
		max_limit := page * PERPAGE
		mod.Limit(min_limit, max_limit)
	}
	mod.Order("id desc")
	resp, err := mod.All()
	if err != nil {
		return res, err
	}

	//是否需要对数据解析
	list := resp.List()
	for _, data := range list {
		data["status"] = devops.StatusList()[gconv.Int(data["status"])]
	}
	res["total"] = total_num
	res["list"] = list
	res["page"] = page

	return res, nil
}

func (r *ReleaseService) Edit(data *releasestruct.EditReleaseParam) error {

	//开启事务
	if tx, err := r.Db.Begin(); err == nil {
		//写入主干数据
		_, err := tx.Table("devops_release").Data(g.Map{
			"title":      data.Title,
			"author":     data.Author,
			"content":    data.Content,
			"updated_at": gtime.Datetime(),
		}).Where("id", data.Id).Update()
		if err != nil {
			tx.Rollback()
			return err
		}

		//修改sql表
		for _, s := range *data.Sql {

			_, err = tx.Table("devops_sql").Data(g.Map{
				"db":         s.Db,
				"crud":       s.Crud,
				"remark":     s.Remark,
				"updated_at": gtime.Datetime(),
			}).Where("id", s.Id).Update()
		}
		if err != nil {
			tx.Rollback()
			return err
		}
		//修改package
		for _, p := range *data.Package {
			_, err = tx.Table("devops_package").Data(g.Map{
				"name":       p.Name,
				"branch":     p.Branch,
				"updated_at": gtime.Datetime(),
			}).Where("id", p.Id).Update()
		}
		if err != nil {
			tx.Rollback()
			return err
		}

		//写入service
		for _, s := range *data.Service {
			_, err = tx.Table("devops_service").Data(g.Map{
				"name":       s.Name,
				"branch":     s.Branch,
				"version":    s.Version,
				"env":        gconv.String(s.Env),
				"script":     gconv.String(s.Script),
				"updated_at": gtime.Datetime(),
			}).Where("id", s.Id).Update()
		}
		if err != nil {
			tx.Rollback()
			return err
		}

		tx.Commit()
	}
	return nil

}

func (r *ReleaseService) Export(release_id string) (*releasestruct.OnlineInfo, error) {
	resp := releasestruct.NewOnlineInfo()
	resp.Id = gconv.Int(release_id)
	//info
	info, err := r.Db.Table("devops_release").Where("id=?", release_id).One()
	if err != nil {
		return resp, err
	}
	res := info.Map()
	if content, ok := res["content"]; ok {
		resp.Content = gconv.String(content)
	}
	if author, ok := res["author"]; ok {
		resp.Author = gconv.String(author)
	}

	//sql
	sql, err := r.Db.Table("devops_sql").Where("release_id=?", release_id).All()
	if err != nil {
		return resp, err
	}
	sql_list := sql.List()
	s_list := []*releasestruct.OnlineSql{}
	for _, item := range sql_list {
		s_list = append(s_list, &releasestruct.OnlineSql{
			Db:     gconv.String(item["db"]),
			Crud:   gconv.String(item["crud"]),
			Remark: gconv.String(item["remark"]),
		})
	}
	resp.Sql = s_list

	//package
	pg, err := r.Db.Table("devops_package").Where("release_id=?", release_id).All()
	if err != nil {
		return resp, err
	}
	package_list := pg.List()
	p_list := []*releasestruct.OnlinePackage{}
	for _, item := range package_list {
		p_list = append(p_list, &releasestruct.OnlinePackage{
			PackageName: gconv.String(item["name"]),
			BranchName:  gconv.String(item["branch"]),
		})
	}
	resp.Package = p_list
	//service
	service, err := r.Db.Table("devops_service").Where("release_id=?", release_id).All()
	if err != nil {
		return resp, err
	}
	service_list := service.List()
	online_s := []*releasestruct.OnlineService{}
	for _, item := range service_list {
		env_list := g.List{}
		script_list := g.List{}
		json.Unmarshal([]byte(gconv.String(item["env"])), &env_list)
		json.Unmarshal([]byte(gconv.String(item["script"])), &script_list)
		rsp_env := []*releasestruct.Env{}
		for _, env_map := range env_list{
			if g.IsEmpty(env_map["handle"]) && g.IsEmpty(env_map["key"]) && g.IsEmpty(env_map["value"]) {
				continue
			}
			rsp_env = append(rsp_env, &releasestruct.Env{
				Handle: gconv.String(env_map["handle"]),
				Key:    gconv.String(env_map["key"]),
				Value:  gconv.String(env_map["value"]),
			})
		}

		rsp_script := []*releasestruct.Script{}
		for _, script_map := range script_list{
			if g.IsEmpty(script_map["cmd"]) && g.IsEmpty(script_map["remark"]) {
				continue
			}
			rsp_script = append(rsp_script, &releasestruct.Script{
				Cmd:    gconv.String(script_map["cmd"]),
				Remark: gconv.String(script_map["remark"]),
			})
		}
		online_s = append(online_s, &releasestruct.OnlineService{
			ServiceName: gconv.String(item["name"]),
			Version:     "v1",
			BranchName:  gconv.String(item["branch"]),
			Env:         rsp_env,
			Script:      rsp_script,
		})
	}
	resp.Service = online_s

	return resp, err
}

//返回一些通用信息
func (r *ReleaseService) CommonInfo() (g.Map, error) {
	res := make(g.Map)

	//数据库名

	return res, nil

}
