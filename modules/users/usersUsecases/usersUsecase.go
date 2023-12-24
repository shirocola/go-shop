package usersUsecases

import (
	"github.com/Rayato159/kawaii-shop-tutorial/modules/users/usersRepositories"
	"github.com/shirocola/go-shop/config"
)

type IUsersUsecases interface {
}

type usersUsecases struct {
	cfg             config.IConfig
	usersRepository usersRepositories.IUsersRepository
}

func UserUsecases(cfg config.IConfig, usersRepository usersRepositories.IUsersRepository) IUsersUsecases {
	return &usersUsecases{
		cfg:             cfg,
		usersRepository: usersRepository,
	}
}
