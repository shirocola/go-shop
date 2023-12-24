package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/shirocola/go-shop/config"
)

type IServer interface {
	Start()
}

type server struct {
	app *fiber.App
	cfg config.IConfig
	db  *sqlx.DB
}

func NewServer(cfg config.IConfig, db *sqlx.DB) IServer {
	return &server{
		cfg: cfg,
		db:  db,
		app: fiber.New(fiber.Config{
			AppName:      cfg.App().Name(),
			BodyLimit:    cfg.App().BodyLimit(),
			ReadTimeout:  cfg.App().ReadTimeout(),
			WriteTimeout: cfg.App().WriteTimeout(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
	}
}

func (s *server) Start() {
	//middlewares
	middlewares := InitMiddlewares(s)
	s.app.Use(middlewares.Cors())

	//modules
	v1 := s.app.Group("v1")
	modules := InitModule(v1, s, middlewares)

	modules.MonitorModule()

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	// listen to host:port
	log.Printf("server is starting on %v", s.cfg.App().Url())
	s.app.Listen(s.cfg.App().Url())
}
