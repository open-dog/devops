package mongo

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"gopkg.in/mgo.v2"
	"log"
	"sync"
	"time"
)

var (
	once sync.Once
	m *mgo.Session
)

//返回单例的
func Mdb() *mgo.Session {
	once.Do(func() {
		m = getSession()
	})

	return m
}

//先简单处理
func getSession() *mgo.Session {
	//获取对应的配置文件
	config := g.Cfg().GetMap("mongodb.default.0")

	dialInfo := &mgo.DialInfo{
		Addrs:          []string{gconv.String(config["address"])},
		Timeout:        time.Second * time.Duration(gconv.Int64(config["timeout"])),
		Database:       gconv.String(config["database"]),
		Source:         gconv.String(config["source"]),
		Username:       gconv.String(config["username"]),
		Password:       gconv.String(config["password"]),
		PoolLimit:      gconv.Int(config["poollimit"]),
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}

	return session
}

