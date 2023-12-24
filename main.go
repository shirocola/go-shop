package main

import (
	"os"

	"github.com/shirocola/go-shop/config"
	"github.com/shirocola/go-shop/modules/servers"
	"github.com/shirocola/go-shop/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnenct(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()

}
