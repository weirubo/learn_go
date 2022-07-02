package http

import (
	// "errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/weirubo/learn_go/lesson41/domain"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(echo *echo.Echo, userUsecase domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: userUsecase,
	}
	echo.GET("/user/:id", handler.GetUserById)
}

func (u *UserHandler) GetUserById(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	id := int64(idP)
	ctx := c.Request().Context()
	user := &domain.User{
		Id: id,
	}
	err = u.UserUsecase.GetUserById(ctx, user)
	if err != nil {
		// err = errors.New("UserUsecase error")
		err = errors.Wrap(err, "UserUsecase error")
		fmt.Printf("UserHandler || GetUserById() || uid=%v || err=%+v\n", id, err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}
