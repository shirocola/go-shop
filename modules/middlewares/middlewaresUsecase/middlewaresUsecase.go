package middlewaresUsecase

import (
	"github.com/shirocola/go-shop/modules/middlewares/middlewaresRepository"
)

type IMiddlewaresUsecase interface {
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepository.IMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepository.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}
