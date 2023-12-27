package usersUsecases

import (
	"github.com/Rayato159/kawaii-shop-tutorial/modules/users/usersRepositories"
	"github.com/shirocola/go-shop/config"
	"github.com/shirocola/go-shop/modules/users"
)

type IUsersUsecases interface {
	InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error)
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

func (u *usersUsecases) InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error) {
	// Hashing password
	if err := req.BcryptHasing(); err != nil {
		return nil, err
	}

	// Insert user
	result, err := u.usersRepository.InsertUser(req, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}
