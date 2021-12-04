package lesson32

import (
	"log"
	"testing"
	"xorm.io/xorm"
)

type DBOrm interface {
	Insert(param ...interface{})
}

type XormDB struct {
	db *xorm.Session
}

func (x *XormDB) Insert(param ...interface{}) {
	_, err := x.db.Insert(param)
	if err != nil {
		log.Println(err)
	}
}

// type GormDB struct {
// 	db *gorm.DB
// }
//
// func (g *GormDB) Insert(param ...interface{}) {
// 	g.db.Create(param)
// }

type User struct {
	orm  DBOrm
	Id   int64
	Name string
}

func (u *User) DB() DBOrm {
	u.orm = new(XormDB) // 数据库实例注入接口
	// u.orm = new(GormDB)
	return u.orm
}

func TestOrm(t *testing.T) {
	user1 := new(User)
	user1.Name = "lucy"
	user1.DB().Insert(user1)
}
