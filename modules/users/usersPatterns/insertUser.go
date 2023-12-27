package usersPatterns

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	users "github.com/shirocola/go-shop/modules"
)

type IInsertUser interface {
	Customer() (IInsertUser, error)
	Admin() (IInsertUser, error)
	Result() (*users.UserPassport, error)
}

type userRequest struct {
	id  string
	req *users.UserRegisterReq
	db  *sqlx.DB
}

type customer struct {
	*userRequest
}

type admin struct {
	*userRequest
}

func InsertUser(db *sqlx.DB, req *users.UserRegisterReq, isAdmin bool) IInsertUser {
	if isAdmin {
		return NewAdmin(db, req)
	}
	return NewCustomer(db, req)

}

func NewCustomer(db *sqlx.DB, req *users.UserRegisterReq) IInsertUser {
	return &customer{
		userRequest: &userRequest{
			req: req,
			db:  db,
		},
	}
}

func NewAdmin(db *sqlx.DB, req *users.UserRegisterReq) IInsertUser {
	return &admin{
		userRequest: &userRequest{
			req: req,
			db:  db,
		},
	}

}

func (f *userRequest) Customer() (IInsertUser, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	query := `
	INSERT INTO users (
		email, 
		password, 
		username, 
		role_id
	)
	VALUES (
		($1, $2, $3, 1)
	RETURNING id;`

	if err := f.db.QueryRowContext(
		ctx,
		query,
		f.req.Email,
		f.req.Password,
		f.req.Username,
	).Scan(&f.id); err != nil {
		switch err.Error() {
		case "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)":
			return nil, errors.New("email already exists")
		case "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)":
			return nil, errors.New("username already exists")
		default:
			return nil, fmt.Errorf("error inserting user: %w", err)
		}

	}
	return f, nil
}

func (f *userRequest) Admin() (IInsertUser, error) {
	// ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancle()

	return nil, nil
}

func (f *userRequest) Result() (*users.UserPassport, error) {
	query := `
	SELECT 
		json_build_object(
			'user', "t"
			'token', NULL
		)
	FROM (
		SELECT
			"u"."id",
			"u"."email",
			"u"."username",
			"u"."role_id",
		FROM "users" "u"
		WHERE "u"."id" = $1
	) AS "t";`

	data := make([]byte, 0)
	if err := f.db.Get(&data, query, f.id); err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	user := new(users.UserPassport)
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("error unmarshaling user: %w", err)
	}
	return user, nil
}
