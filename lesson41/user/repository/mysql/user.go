package mysql

import (
	"context"
	"fmt"
	"github.com/weirubo/learn_go/lesson41/domain"
	"xorm.io/xorm"
)

type mysqlUserRepository struct {
	DB *xorm.Engine
}

func NewMysqlUserRepository(db *xorm.Engine) domain.UserRepository {
	return &mysqlUserRepository{
		DB: db,
	}
}

func (m *mysqlUserRepository) GetUserById(ctx context.Context, user *domain.User) (err error) {
	_, err = m.DB.Get(user)
	fmt.Printf("mysqlUserRepository || GetUserById() || uid=%v || err=%v\n", user.Id, err)
	return
}
