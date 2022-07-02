package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	_userHttpDelivery "github.com/weirubo/learn_go/lesson41/user/delivery/http"
	_userRepo "github.com/weirubo/learn_go/lesson41/user/repository/mysql"
	_userUsecase "github.com/weirubo/learn_go/lesson41/user/usecase"
	"xorm.io/xorm"
)

func main() {
	db, err := xorm.NewEngine("mysql", "root:mysqlpw123@tcp(127.0.0.1:55000)/test")
	if err != nil {
		fmt.Printf("main || xorm || err=%v\n", err)
		return
	}
	defer func() {
		err := db.Close()
		if err != nil {
			return
		}
	}()
	e := echo.New()
	userRepo := _userRepo.NewMysqlUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	_userHttpDelivery.NewUserHandler(e, userUsecase)
	err = e.Start(":1323")
	if err != nil {
		fmt.Printf("main || echo start || err=%v\n", err)
		return
	}
}
