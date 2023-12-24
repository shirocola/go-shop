package middlewaresHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/shirocola/go-shop/config"
	"github.com/shirocola/go-shop/modules/entities"
	"github.com/shirocola/go-shop/modules/middlewares/middlewaresUsecase"
)

type middlewaresHandlerErrCode string

const (
	routerCheckErr middlewaresHandlerErrCode = "router-001"
)

type IMiddlewaresHandler interface {
	Cors() fiber.Handler
	RouterCheck() fiber.Handler
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

func (h *middlewaresHandler) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return entities.NewResponse(c).Error(
			fiber.ErrNotFound.Code,
			string(routerCheckErr),
			"router not found",
		).Res()
	}
}
