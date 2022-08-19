package httpsvc

import (
	"github.com/bagasss3/go-weight/internal/model"
	"github.com/labstack/echo/v4"
)

type Service struct {
	group            *echo.Group
	weightController model.WeightController
	weightRepository model.WeightRepository
}

func RouteService(
	group *echo.Group,
	weightController model.WeightController,
	weightRepository model.WeightRepository,
) {
	srv := &Service{
		group:            group,
		weightController: weightController,
		weightRepository: weightRepository,
	}
	srv.initRoutes()
}

func (s *Service) initRoutes() {
	s.group.POST("/weight", s.handleCreateWeight())
	s.group.GET("/weight", s.handleGetAllWeights())
	s.group.GET("/weight/:id", s.handleShowWeight())
	s.group.PUT("/weight/:id", s.handleUpdateWeight())
	s.group.DELETE("/weight/:id", s.handleDeleteWeight())
}
