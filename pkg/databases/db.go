package databases

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/shirocola/go-shop/config"
)

func DbConnenct(cfg config.IDbConfig) *sqlx.DB {
	//Connect
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("connect to db failed: %v", err)
	}
	db.DB.SetMaxOpenConns(cfg.MaxOpenConns())
	return db
}
