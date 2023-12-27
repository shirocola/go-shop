package userrepositories

import (
	"github.com/jmoiron/sqlx"
	users "github.com/shirocola/go-shop/modules"
	"github.com/shirocola/go-shop/modules/users/usersPatterns"
)

type IUsersRepository interface {
}

type usersRepository struct {
	db *sqlx.DB
}

func UsersRepository(db *sqlx.DB) IUsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error) {
	result := usersPatterns.InsertUser(r.db, req, isAdmin)

	var err error
	if isAdmin {
		result, err = result.Admin()
		if err != nil {
			return nil, err
		}
	} else {
		result, err = result.Customer()
		if err != nil {
			return nil, err
		}
	}

	// Get reult from inserting
	user, err := result.Result()
	if err != nil {
		return nil, err
	}

	return user, nil

}
