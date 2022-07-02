package usecase

import (
	"context"
	// "errors"
	"fmt"
	"github.com/pkg/errors"
	"github.com/weirubo/learn_go/lesson41/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) GetUserById(ctx context.Context, user *domain.User) (err error) {
	if user.Id == 0 {
		err = errors.New("invalid request parameter")
	}
	err = u.userRepo.GetUserById(ctx, user)
	fmt.Printf("userUsecase || GetUserById() || uid=%v || err=%v\n", user.Id, err)
	return
}
