package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shirocola/go-shop/modules/middlewares/middlewaresHandler"
	"github.com/shirocola/go-shop/modules/middlewares/middlewaresRepository"
	"github.com/shirocola/go-shop/modules/middlewares/middlewaresUsecase"
	"github.com/shirocola/go-shop/modules/monitor/monitorHandlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type ModuleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandler.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandler.IMiddlewaresHandler) IModuleFactory {
	return &ModuleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandler.IMiddlewaresHandler {
	repository := middlewaresRepository.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecase.MiddlewaresUsecase(repository)
	return middlewaresHandler.MiddlewaresHandler(usecase)

}

func (m *ModuleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HealthCheck)
}
