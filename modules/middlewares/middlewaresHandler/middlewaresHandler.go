package middlewaresHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/shirocola/go-shop/config"
	"github.com/shirocola/go-shop/modules/middlewares/middlewaresUsecase"
)

type IMiddlewaresHandler interface {
	Cors() fiber.Handler
}

type middlewaresHandler struct {
	cfg                config.IConfig
	middlewaresUsecase middlewaresUsecase.IMiddlewaresUsecase
}

func MiddlewaresHandler(middlewaresUsecase middlewaresUsecase.IMiddlewaresUsecase) IMiddlewaresHandler {
	return &middlewaresHandler{
		middlewaresUsecase: middlewaresUsecase,
	}
}

func (h *middlewaresHandler) Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}
