package middlewaresHandler

import (
	"github.com/shirocola/go-shop/config"
	"github.com/shirocola/go-shop/modules/middlewares/middlewaresUsecase"
)

type IMiddlewaresHandler interface {
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
