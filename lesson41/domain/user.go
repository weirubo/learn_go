package domain

import "context"

type User struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Created int    `json:"created"`
	Updated int    `json:"updated"`
}

type UserUsecase interface {
	GetUserById(ctx context.Context, user *User) error
}

type UserRepository interface {
	GetUserById(ctx context.Context, user *User) error
}
