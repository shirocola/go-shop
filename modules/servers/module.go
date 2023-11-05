package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shirocola/go-shop/modules/monitor/monitorHandlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type ModuleFactory struct {
	r fiber.Router
	s *server
}

func InitModule(r fiber.Router, s *server) IModuleFactory {
	return &ModuleFactory{
		r: r,
		s: s,
	}
}

func (m *ModuleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HealthCheck)
}
