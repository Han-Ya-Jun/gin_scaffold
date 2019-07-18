package models

import (
	"github.com/Han-Ya-Jun/gin_scaffold/pkg/setting"
	"github.com/globalsign/mgo"
	"log"
	"strings"
)

/*
* @Author:hanyajun
* @Date:2019/5/23 23:54
* @Name:mongo的基本操作
* @Function:
 */
var globalS *mgo.Session

const (
	db = "gin_scaffold"
)

func SetUp() {
	hostList := strings.Split(setting.MongoSetting.Host, ";")

	dialInfo := &mgo.DialInfo{
		Addrs:    hostList,
		Source:   setting.MongoSetting.Source,
		Username: setting.MongoSetting.User,
		Password: setting.MongoSetting.Password,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("create session error ", err)
		panic(err)
	}
	s.SetMode(mgo.Monotonic, true)
	globalS = s
}

func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

func Inserts(db, collection string, docs []interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}
func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func Update(db, collection string, query, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(query)
}

func IsEmpty(db, collection string) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func Count(db, collection string, query interface{}) (int, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Count()
}
