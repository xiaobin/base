package mymongo

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"time"
)

var session *mgo.Session

// Conn return mongodb session.
func Conn() *mgo.Session {
	return session.Copy()
}
func Print(){
	beego.Debug("go,mongodb")
}
/*
func Close() {
	session.Close()
}
*/

func init() {
	beego.Debug("初始化mongodb")
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{beego.AppConfig.String("mongodb.ip")},
		Direct:   false,
		Timeout:  time.Second * 1,
		Database: beego.AppConfig.String("mongodb.database"),
		//Source:    "admin",
		Username:  beego.AppConfig.String("mongodb.username"),
		Password:  beego.AppConfig.String("mongodb.password"),
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	session = s
	session.SetMode(mgo.Monotonic, true)
}
