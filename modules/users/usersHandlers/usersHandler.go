package usersHandlers

import (
	"github.com/shirocola/go-shop/config"
	"github.com/shirocola/go-shop/modules/users/usersUsecases"
)

type IUsersHandler interface {
}

type usersHandler struct {
	cfg           config.IConfig
	usersUsecases usersUsecases.IUsersUsecases
}

func UsersHandler() IUsersHandler {
	return &usersHandler{}
}
